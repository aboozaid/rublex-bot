package telegram

import "os"

type Config struct {
	TelegramApiAppID   string
	TelegramApiAppHash string
	TelegramBotToken   string
}

func GetConfig() Config {
	return Config{
		TelegramApiAppID:   os.Getenv("TELEGRAM_API_APP_ID"),
		TelegramApiAppHash: os.Getenv("TELEGRAM_API_APP_HASH"),
		TelegramBotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}
