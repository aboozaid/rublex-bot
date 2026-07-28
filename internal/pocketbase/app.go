package pocketbase

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	tgAuthPlugin "github.com/aboozaid/pocketbase-plugin-telegram-auth"
)

func New() *pocketbase.PocketBase {
	pbCfg := pocketbase.Config{}
	config := GetConfig()

	if config.PocketbaseEncryptionKey != "" {
		if !ValidateEncryptionKey(config.PocketbaseEncryptionKey) {
			slog.Error("POCKETBASE_ENCRYPTION_KEY msut be 32-character hexadecimal string (generated with 'openssl rand -hex 16')")
			os.Exit(1)
		}
		pbCfg.DefaultEncryptionEnv = "POCKETBASE_ENCRYPTION_KEY"
	}

	app := pocketbase.NewWithConfig(pbCfg)

	if config.Environment == "development" {
		migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{Automigrate: true})
	}

	tgAuthPlugin.MustRegister(app, &tgAuthPlugin.Options{BotToken: config.PocketbaseOauthTGBotToken, CollectionKey: "users"})

	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		settings := app.Settings()
		settings.RateLimits.Enabled = config.RateLimitEnabled == "true"
		if settings.RateLimits.Enabled && config.RateLimitRules != "" {
			var rules []core.RateLimitRule
			if err := json.Unmarshal([]byte(config.RateLimitRules), &rules); err != nil {
				slog.Warn("Failed to parse RATE_LIMIT_RULES, using defaults", "error", err)
			} else {
				settings.RateLimits.Rules = rules
			}
		}
		return e.Next()
	})

	return app
}
