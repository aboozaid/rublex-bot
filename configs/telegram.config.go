package configs

import (
	"p2ptrader/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TelegramConfig struct {
	TelegramApiAppID   string
	TelegramApiAppHash string
	TelegramBotToken   string
	TelegramAdminID    string
}

func LoadTelegramConfig() TelegramConfig {
	return TelegramConfig{
		TelegramApiAppID:   utils.Getenv("TELEGRAM_API_APP_ID"),
		TelegramApiAppHash: utils.Getenv("TELEGRAM_API_APP_HASH"),
		TelegramBotToken:   utils.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramAdminID:    utils.Getenv("TELEGRAM_ADMIN_ID"),
	}
}

func (c TelegramConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.TelegramApiAppID, validation.Required),
		validation.Field(&c.TelegramApiAppHash, validation.Required),
		validation.Field(&c.TelegramBotToken, validation.Required),
		validation.Field(&c.TelegramAdminID, validation.Required),
	)
}
