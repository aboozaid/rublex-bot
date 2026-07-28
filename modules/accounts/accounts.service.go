package accounts

import (
	"context"
	"p2ptrader/modules/accounts/dto"
	"p2ptrader/modules/accounts/models"
	"p2ptrader/modules/database"
	exchangesDto "p2ptrader/modules/exchanges/dto"
	"p2ptrader/modules/kms"

	"github.com/pocketbase/pocketbase/core"
)

type Service interface {
	CreateAccount(ctx context.Context, userID string, createAccountParams dto.CreateAccountRequestParams) (string, error)
	GetAccountsByTelegramID(ctx context.Context, telegramID string) ([]dto.Account, error)
	GetExchangeKeysByAccountID(ctx context.Context, accountID string) (dto.GetEncryptedExchangeKeysResponseParams, error)
	GetUserIDByAccountID(ctx context.Context, accountID string) (string, error)
	//IsTgGroupExist(ctx context.Context, tgGroupID string) (bool, error)

	//SetTgGroupByAccountID(ctx context.Context, accountID, tgGroupID string) error
}

type service struct {
	repository repo
	dbService  database.Service
	kmsService kms.Service
}

func newService(app core.App, dbService database.Service, kmsService kms.Service) Service {
	repository := newRepository(app)

	return service{repository, dbService, kmsService}
}

func (s service) CreateAccount(ctx context.Context, userID string, createAccountParams dto.CreateAccountRequestParams) (string, error) {
	encryptedApiKey, err := s.kmsService.Encrypt(createAccountParams.ApiKey)
	if err != nil {
		return "", err
	}
	encryptedApiSecret, err := s.kmsService.Encrypt(createAccountParams.ApiSecret)
	if err != nil {
		return "", err
	}
	accountID, err := s.repository.CreateAccount(ctx, func(account *models.Account) error {
		account.SetName(createAccountParams.Name)
		account.SetApiKey(encryptedApiKey)
		account.SetApiSecret(encryptedApiSecret)
		account.SetIsActive(true)
		if createAccountParams.Nickname != nil {
			account.SetNickname(*createAccountParams.Nickname)
		}
		if createAccountParams.CountryCode != nil {
			account.SetCountryCode(*createAccountParams.CountryCode)
		}
		/*json, err := json.Marshal(createAccountParams.Metadata)
		if err != nil {
			return err
		}
		account.SetMetadata(string(json))*/

		account.SetUser(userID)
		account.SetExchange(createAccountParams.ExchangeID)

		apiKeyHashed := s.kmsService.Hash(createAccountParams.ApiKey)
		account.SetApiKeyHash(apiKeyHashed)
		return nil
	})
	return accountID, err
}

func (s service) GetAccountsByTelegramID(ctx context.Context, telegramID string) ([]dto.Account, error) {
	dbAccounts, err := s.repository.GetAccountsByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}
	accounts := make([]dto.Account, len(dbAccounts))
	for i, a := range dbAccounts {
		account := dto.Account{ID: a.Id, Name: a.Name(), Nickname: a.Nickname(), Exchange: a.Exchange().Name()}
		accounts[i] = account
	}
	return accounts, nil
}

func (s service) GetUserIDByAccountID(ctx context.Context, accountID string) (string, error) {
	account, err := s.repository.GetAccountByID(ctx, accountID)
	if err != nil {
		return "", err
	}
	return account.UserID(), nil
}

/*func (s service) IsTgGroupExist(ctx context.Context, tgGroupID string) (bool, error) {
	total, err := s.repository.GetTotalAccountsByTgGroupID(ctx, tgGroupID)
	if err != nil {
		return false, err
	}
	return total > 0, nil
}*/

/*func (s service) SetTgGroupByAccountID(ctx context.Context, accountID, tgGroupID string) error {
	err := s.repository.UpdateAccount(ctx, accountID, func(a *models.Account) error {
		a.SetTgGroupID(tgGroupID)
		return nil
	})
	return err
}*/

func (s service) GetExchangeKeysByAccountID(ctx context.Context, accountID string) (dto.GetEncryptedExchangeKeysResponseParams, error) {
	account, err := s.repository.GetAccountByID(ctx, accountID)
	if err != nil {
		return dto.GetEncryptedExchangeKeysResponseParams{}, err
	}
	exchange := account.Exchange()
	exchangeKeys := dto.GetEncryptedExchangeKeysResponseParams{
		Exchange:  exchangesDto.Exchange{ID: exchange.Id, Name: exchange.Name()},
		ApiKey:    account.ApiKey(),
		ApiSecret: account.ApiSecret(),
	}
	return exchangeKeys, nil
}
