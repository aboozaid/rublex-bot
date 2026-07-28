package binance

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"p2ptrader/modules/accounts"
	accountsDto "p2ptrader/modules/accounts/dto"
	"p2ptrader/modules/binance/dto"
	"p2ptrader/modules/cache"
	"p2ptrader/modules/http"
	"p2ptrader/utils"

	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	VerifyAccount(ctx context.Context, userID string, exchangeID string, verifyAccountRequestParams dto.VerifyAccountRequestParams) (dto.VerifyAccountResponseParams, error)
}

type binanceKeys struct {
	ApiKey    string
	ApiSecret string
}

type ctxBinanceKey string

const ctxBinanceKeys ctxBinanceKey = "ctx_binance_keys"

var (
	apiKeyIdentifier    = "apiKey"
	apiSecretIdentifier = "apiSecret"
)

type service struct {
	repository     repo
	cacheService   cache.Service
	httpService    http.Service
	accountService accounts.Service
}

func newService(app core.App, cacheService cache.Service, httpService http.Service, accountService accounts.Service) Service {
	repository := newRepository(app)
	return service{repository, cacheService, httpService, accountService}
}

func (s service) VerifyAccount(ctx context.Context, userID string, exchangeID string, verifyAccountRequestParams dto.VerifyAccountRequestParams) (dto.VerifyAccountResponseParams, error) {
	var binanceUser dto.BinanceResponse[dto.BinanceUser]
	cctx := contextWithBinanceKeys(ctx, verifyAccountRequestParams.ApiKey, verifyAccountRequestParams.ApiSecret)
	_, err := s.httpService.Client.R().
		SetContext(cctx).
		SetSuccessResult(&binanceUser).
		Post("/sapi/v1/c2c/user/baseDetail")
	if err != nil {
		return dto.VerifyAccountResponseParams{}, err
	}

	createAccountParams := accountsDto.CreateAccountRequestParams{
		ExchangeID:  exchangeID,
		Name:        binanceUser.Data.KycFullname,
		ApiKey:      verifyAccountRequestParams.ApiKey,
		ApiSecret:   verifyAccountRequestParams.ApiSecret,
		CountryCode: &binanceUser.Data.CountryCode,
		//Metadata:    metadata,
	}
	if binanceUser.Data.IsNicknameExists {
		createAccountParams.Nickname = &binanceUser.Data.Nickname
	}
	accountID, err := s.accountService.CreateAccount(ctx, userID, createAccountParams)
	if err != nil {
		fmt.Println(err.Error())
		return dto.VerifyAccountResponseParams{}, err
	}

	s.cacheService.Set(accountID+"|"+apiKeyIdentifier, verifyAccountRequestParams.ApiKey)
	s.cacheService.Set(accountID+"|"+apiSecretIdentifier, verifyAccountRequestParams.ApiSecret)

	tgBotName := fmt.Sprintf("RublexBot <> %s <> Binance", binanceUser.Data.Nickname)
	tgBotUsername := utils.GenerateTgUsername("rublex")

	return dto.VerifyAccountResponseParams{AccountID: accountID, TelegramBotName: tgBotName, TelegramBotUsername: tgBotUsername}, nil
}

func contextWithBinanceKeys(ctx context.Context, apiKey, apiSecret string) context.Context {
	keys := binanceKeys{ApiKey: apiKey, ApiSecret: apiSecret}
	context := context.WithValue(ctx, ctxBinanceKeys, keys)
	return context
}

func getBinanceKeys(ctx context.Context) (binanceKeys, bool) {
	keys, ok := ctx.Value(ctxBinanceKeys).(binanceKeys)
	return keys, ok
}

func createSignature(payload, apiSecret string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(payload))
	sig := hex.EncodeToString(mac.Sum(nil))

	return sig
}
