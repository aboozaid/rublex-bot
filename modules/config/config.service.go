package config

import (
	"log/slog"
	"os"
	"p2ptrader/configs"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service interface {
	App() configs.AppConfig
	Telegram() configs.TelegramConfig
	Binance() configs.BinanceConfig
}

type service struct {
	app      configs.AppConfig
	telegram configs.TelegramConfig
	binance  configs.BinanceConfig
}

func newService() Service {
	service := service{app: configs.LoadAppConfig(), telegram: configs.LoadTelegramConfig(), binance: configs.LoadBinanceConfig()}
	if err := service.Validate(); err != nil {
		slog.Error("Unable to load configs", "error", err)
		os.Exit(1)
	}
	return service
}

func (s service) App() configs.AppConfig {
	return s.app
}

func (s service) Telegram() configs.TelegramConfig {
	return s.telegram
}

func (s service) Binance() configs.BinanceConfig {
	return s.binance
}

func (s service) Validate() error {
	return validation.ValidateStruct(&s, validation.Field(&s.app), validation.Field(&s.telegram), validation.Field(&s.binance))
}
