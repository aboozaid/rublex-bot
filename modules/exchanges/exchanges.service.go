package exchanges

import (
	"context"
	"p2ptrader/modules/binance"
	"p2ptrader/modules/database"
	"p2ptrader/modules/exchanges/dto"
	"p2ptrader/modules/exchanges/models"

	binanceDto "p2ptrader/modules/binance/dto"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	VerifyAccount(ctx context.Context, userID string, exchangeID string, verifyAccountParams dto.VerifyAccountRequestParams) (dto.VerifyAccountResponseParams, error)
}

type service struct {
	repository     repo
	dbService      database.Service
	binanceService binance.Service
}

func newService(app core.App, dbService database.Service, binanceService binance.Service) Service {
	repository := newRepository(app)
	return service{repository, dbService, binanceService}
}

func (s service) VerifyAccount(ctx context.Context, userID string, exchangeID string, verifyAccountParams dto.VerifyAccountRequestParams) (dto.VerifyAccountResponseParams, error) {
	exhange, err := s.repository.GetExchangeByID(ctx, exchangeID)
	if err != nil {
		return dto.VerifyAccountResponseParams{}, err
	}
	switch exhange.ExchangeType() {
	case models.Binance:
		verifyBinanceAccountParams := binanceDto.VerifyAccountRequestParams{ApiKey: verifyAccountParams.ApiKey, ApiSecret: verifyAccountParams.ApiSecret}
		binanceAccount, err := s.binanceService.VerifyAccount(ctx, userID, exchangeID, verifyBinanceAccountParams)
		if err != nil {
			return dto.VerifyAccountResponseParams{}, err
		}
		return dto.VerifyAccountResponseParams{AccountID: binanceAccount.AccountID, TelegramBotName: binanceAccount.TelegramBotName, TelegramBotUsername: binanceAccount.TelegramBotUsername}, nil
	default:
		return dto.VerifyAccountResponseParams{}, apis.NewInternalServerError("unsupported exchange", nil)
	}
}
