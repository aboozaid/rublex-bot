package binance

import (
	"context"
	"fmt"
	"p2ptrader/modules/database"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type repo interface {
	FindUserByTelegramID(ctx context.Context, telegramID string) (*core.Record, error)
	UpdateUserVerificationStatus(ctx context.Context, userID string, isVerified bool) error
}

type repository struct {
	db core.App
}

const COLLECTION_NAME = "users"

func newRepository(db core.App) repo {
	return repository{db}
}

func (r repository) FindUserByTelegramID(ctx context.Context, telegramID string) (*core.Record, error) {
	user, err := database.GetDB(ctx, r.db).FindFirstRecordByData(COLLECTION_NAME, "telegram_id", telegramID)
	return user, err
}

func (r repository) UpdateUserVerificationStatus(ctx context.Context, userID string, isVerified bool) error {
	query := database.GetDB(ctx, r.db).DB().Update("users", dbx.Params{"verified": isVerified}, dbx.NewExp(fmt.Sprintf("id = '%s'", userID)))
	_, err := query.Execute()
	return err
}
