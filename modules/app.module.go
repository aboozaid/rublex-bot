package modules

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"p2ptrader/common"
	"p2ptrader/modules/accounts"
	"p2ptrader/modules/binance"
	"p2ptrader/modules/cache"
	"p2ptrader/modules/config"
	"p2ptrader/modules/exchanges"
	"p2ptrader/modules/invitations"
	tg "p2ptrader/modules/telegram"
	telegramgroups "p2ptrader/modules/telegram-groups"
	"p2ptrader/modules/users"
	"p2ptrader/utils"
	"strconv"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	tgAuthPlugin "github.com/aboozaid/pocketbase-plugin-telegram-auth"
)

type Module interface {
	Start() error
}

type module struct {
	app        *pocketbase.PocketBase
	tgClient   *telegram.Client
	ctgService config.Service

	modules []any
}

func NewAppModule() Module {
	config := config.NewModule().GetService()

	pbCfg := pocketbase.Config{}
	if config.App().PocketbaseEncryptionKey != "" {
		if !utils.ValidateEncryptionKey(config.App().PocketbaseEncryptionKey) {
			slog.Error("POCKETBASE_ENCRYPTION_KEY msut be 32-character hexadecimal string (generated with 'openssl rand -hex 16')")
			os.Exit(1)
		}
		pbCfg.DefaultEncryptionEnv = "POCKETBASE_ENCRYPTION_KEY"
	}

	telegramCfg := config.Telegram()
	appID, err := strconv.ParseInt(telegramCfg.TelegramApiAppID, 10, 32)
	if err != nil {
		slog.Error("could not parse TELEGRAM_API_APP_ID value", "error", err)
		os.Exit(1)
	}
	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID: int32(appID), 
		AppHash: telegramCfg.TelegramApiAppHash,
		Session: "gogram/session.dat",
		DisableCache: true, // NOTE: We may need to enable this in future
	},
)
	if err != nil {
		slog.Error("could not create telegram client instance", "error", err)
		os.Exit(1)
	}
	commands := []*telegram.BotCommand{
		{
			Command:     "connect",
			Description: "Connect your group to our automation bot",
		},
	}
	client.BotsSetBotCommands(&telegram.BotCommandScopeChatAdmins{}, "", commands)

	app := pocketbase.NewWithConfig(pbCfg)

	users := users.NewModule(app)
	invitations := invitations.NewModule(app, users)
	accounts := accounts.NewModule(app)
	cache := cache.NewModule()
	binance := binance.NewModule(app, cache, accounts)
	exchanges := exchanges.NewModule(app, binance)
	telegramGroups := telegramgroups.NewModule(app)
	telegram := tg.NewModule(accounts, telegramGroups, users)

	modules := []any{users, invitations, accounts, cache, binance, exchanges, telegramGroups, telegram}

	return module{app, client, config, modules}
}

func (m module) Start() error {
	appCfg := m.ctgService.App()
	if appCfg.Environment == "development" {
		migratecmd.MustRegister(m.app, m.app.RootCmd, migratecmd.Config{Automigrate: true})
	}
	botToken := m.ctgService.Telegram().TelegramBotToken
	tgAuthPlugin.MustRegister(m.app, &tgAuthPlugin.Options{BotToken: botToken, CollectionKey: "users"})

	m.app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		settings := m.app.Settings()
		settings.RateLimits.Enabled = appCfg.RateLimitEnabled == "true"
		if settings.RateLimits.Enabled && appCfg.RateLimitRules != "" {
			var rules []core.RateLimitRule
			if err := json.Unmarshal([]byte(appCfg.RateLimitRules), &rules); err != nil {
				slog.Warn("Failed to parse RATE_LIMIT_RULES, using defaults", "error", err)
			} else {
				settings.RateLimits.Rules = rules
			}
		}
		return e.Next()
	})

	m.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		if appCfg.Profile == "just" || appCfg.Profile == "docker" {
			dir, _ := os.Getwd()
			executablePath := fmt.Sprintf("%s/main", dir)
			cmd := exec.Command(executablePath, "superuser", "upsert", appCfg.PocketbaseAdminEmail, appCfg.PocketbaseAdminPassword)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				slog.Error("Superuser upsert failed", "error", err)
			}
		}
		return e.Next()
	})

	m.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		if _, err := m.tgClient.Conn(); err != nil {
			return err
		}
		if err := m.tgClient.LoginBot(botToken); err != nil {
			return err
		}

		go func() {
			m.tgClient.Idle()
		}()

		router := e.Router.Group(appCfg.PocketbaseApiPrefix).
			Bind(apis.RequireAuth())

		for _, mm := range m.modules {
			if v, ok := mm.(common.Controller); ok {
				v.RegisterRoutes(router)
			}
			if v, ok := mm.(common.Handler); ok {
				v.RegisterHandlers(m.tgClient)
			}
			if v, ok := mm.(common.AppLifecycle); ok {
				if err := v.OnAppStart(); err != nil {
					return err
				}
			}
		}

		return e.Next()
	})

	m.app.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		for _, m := range m.modules {
			if v, ok := m.(common.AppLifecycle); ok {
				if err := v.OnAppStop(); err != nil {
					return err
				}
			}
		}
		if m.tgClient.IsConnected() {
			if err := m.tgClient.Stop(); err != nil {
				return err
			}
		}
		return e.Next()
	})

	m.app.OnRecordAuthRequest().BindFunc(func(e *core.RecordAuthRequestEvent) error {
		if e.Record != nil && !e.Record.IsNew() && !e.Record.Verified() {
			return e.UnauthorizedError("UNVERIFIED_USER_NOT_ALLOWED", nil)
		}
		return e.Next()
	})

	return m.app.Start()
}

func Upsert() error {
	app := pocketbase.New()

	return app.Start()
}

func Migrate() error {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{Automigrate: false})

	return app.Start()
}
