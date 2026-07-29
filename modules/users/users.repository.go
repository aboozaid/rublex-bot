package users

import (
	"context"
	"fmt"
	"p2ptrader/modules/database"
	"p2ptrader/modules/users/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type repo interface {
	FindUserByTelegramID(ctx context.Context, telegramID string) (*core.Record, error)
	CreateUser(ctx context.Context, cb func(*models.User) error) (string, error)
	UpdateUserVerificationStatus(ctx context.Context, userID string, isVerified bool) error
}

type repository struct {
	db core.App
}

var collectionKey = "users"

func NewRepository(db core.App) repo {
	return repository{db}
}

func (r repository) FindUserByTelegramID(ctx context.Context, telegramID string) (*core.Record, error) {
	user, err := database.GetDB(ctx, r.db).FindFirstRecordByData(collectionKey, "telegram_id", telegramID)
	return user, err
}

func (r repository) CreateUser(ctx context.Context, cb func(*models.User) error) (string, error) {
	collection, err := r.db.FindCachedCollectionByNameOrId(collectionKey)
	if err != nil {
		return "", err
	}
	user := &models.User{}
	user.SetProxyRecord(core.NewRecord(collection))
	if err := cb(user); err != nil {
		return "", err
	}
	if err := database.GetDB(ctx, r.db).Save(user); err != nil {
		return "", err
	}
	
	return user.Id, nil
}

func (r repository) UpdateUserVerificationStatus(ctx context.Context, userID string, isVerified bool) error {
	query := database.GetDB(ctx, r.db).DB().Update("users", dbx.Params{"verified": isVerified}, dbx.NewExp(fmt.Sprintf("id = '%s'", userID)))
	_, err := query.Execute()
	return err
}
