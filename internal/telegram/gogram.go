package telegram

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/amarnathcjd/gogram/telegram"
)

type Telegram struct {
	client   *telegram.Client
	botToken string
}

func New() Telegram {
	config := GetConfig()
	appID, err := strconv.ParseInt(config.TelegramApiAppID, 10, 32)
	if err != nil {
		slog.Error("could not parse TELEGRAM_API_APP_ID value", "error", err)
		os.Exit(1)
	}
	client, err := telegram.NewClient(telegram.ClientConfig{AppID: int32(appID), AppHash: config.TelegramApiAppHash})
	if err != nil {
		slog.Error("could not create telegram client instance", "error", err)
		os.Exit(1)
	}
	return Telegram{client, config.TelegramBotToken}
}

func (t Telegram) Start() error {
	if _, err := t.client.Conn(); err != nil {
		return err
	}
	if err := t.client.LoginBot(t.botToken); err != nil {
		return err
	}
	go func() {
		t.client.Idle()
	}()
	return nil
}

func (t Telegram) Stop() error {
	if t.client.IsConnected() {
		return t.client.Stop()
	}
	return nil
}
