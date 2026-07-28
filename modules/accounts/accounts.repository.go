package accounts

import (
	"context"
	"fmt"
	"p2ptrader/modules/accounts/models"
	"p2ptrader/modules/database"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

var collectionKey = "accounts"

type repo interface {
	GetAccountByID(ctx context.Context, accountID string) (*models.Account, error)
	GetAccountsByTelegramID(ctx context.Context, telegramID string) ([]*models.Account, error)
	GetTotalAccountsByTgGroupID(ctx context.Context, tgGroupID string) (int, error)
	CreateAccount(ctx context.Context, cb func(*models.Account) error) (string, error)
	UpdateAccount(ctx context.Context, accountID string, cb func(*models.Account) error) error
}

type repository struct {
	db core.App
}

func newRepository(db core.App) repo {
	return repository{db}
}

func (r repository) GetAccountByID(ctx context.Context, accountID string) (*models.Account, error) {
	account := &models.Account{}
	record, err := database.GetDB(ctx, r.db).FindRecordById(collectionKey, accountID)
	/*err := r.db.RecordQuery(collectionKey).
	Select("api_key", "api_secret", "exchanges.id AS exchange.id", "exchange.name AS exchange.name").
	InnerJoin("exchanges", dbx.NewExp("exchanges.id = accounts.exchange")).
	Where(dbx.NewExp("accounts.id = {:id}", dbx.Params{"id": accountID})).
	One(&account)*/
	if err != nil {
		return nil, err
	}
	account.Record = record

	return account, nil
}

func (r repository) GetTotalAccountsByTgGroupID(ctx context.Context, tgGroupID string) (int, error) {
	var totalCount int
	err := database.GetDB(ctx, r.db).DB().
		Select("COUNT(*)").
		From(collectionKey).
		Where(dbx.NewExp("tg_group_id = {:tgGroupID}", dbx.Params{"tgGroupID": tgGroupID})).
		Row(&totalCount)

	return totalCount, err
}

func (r repository) GetAccountsByTelegramID(ctx context.Context, telegramID string) ([]*models.Account, error) {
	records, err := r.db.FindRecordsByFilter(
		collectionKey,
		"user.telegram_id = {:telegramID} && user.verified = true && tg_group_id = ''",
		"",
		0,
		0,
		dbx.Params{"telegramID": telegramID},
	)
	errs := r.db.ExpandRecords(records, []string{"exchange"}, nil)
	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to expand the query: %v", errs)
	}

	if err != nil {
		return nil, err
	}
	accounts := make([]*models.Account, len(records))
	for i, r := range records {
		account := &models.Account{}
		account.SetProxyRecord(r)
		accounts[i] = account
	}

	return accounts, nil
}

func (r repository) CreateAccount(ctx context.Context, cb func(*models.Account) error) (string, error) {
	collection, err := r.db.FindCachedCollectionByNameOrId(collectionKey)
	if err != nil {
		return "", err
	}
	account := &models.Account{}
	account.SetProxyRecord(core.NewRecord(collection))
	if err := cb(account); err != nil {
		return "", err
	}
	if err := database.GetDB(ctx, r.db).Save(account); err != nil {
		return "", err
	}
	return account.Id, nil
}

func (r repository) UpdateAccount(ctx context.Context, accountID string, cb func(*models.Account) error) error {
	account, err := r.GetAccountByID(ctx, accountID)
	if err != nil {
		return err
	}
	if err := cb(account); err != nil {
		return err
	}
	if err := database.GetDB(ctx, r.db).Save(account); err != nil {
		return err
	}
	return nil
}
