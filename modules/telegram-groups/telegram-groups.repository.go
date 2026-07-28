package telegramgroups

import (
	"context"
	"p2ptrader/modules/database"
	"p2ptrader/modules/telegram-groups/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

var collectionKey = "telegram_groups"

type repo interface {
	GetTotalGroupsByGroupID(ctx context.Context, groupID string) (int, error)
	CreateGroup(ctx context.Context, cb func(*models.Group) error) (string, error)
	CreateAccountGroup(ctx context.Context, cb func(*models.AccountGroup) error) error
}

type repository struct {
	db core.App
}

func newRepository(db core.App) repo {
	return repository{db}
}

func (r repository) GetTotalGroupsByGroupID(ctx context.Context, groupID string) (int, error) {
	var totalCount int
	err := database.GetDB(ctx, r.db).DB().
		Select("COUNT(*)").
		From(collectionKey).
		Where(dbx.NewExp("group_id = {:groupID}", dbx.Params{"GroupID": groupID})).
		Row(&totalCount)

	return totalCount, err
}

func (r repository) CreateGroup(ctx context.Context, cb func(*models.Group) error) (string, error) {
	collection, err := r.db.FindCachedCollectionByNameOrId(collectionKey)
	if err != nil {
		return "", err
	}
	group := &models.Group{}
	group.SetProxyRecord(core.NewRecord(collection))
	if err := cb(group); err != nil {
		return "", err
	}
	if err := database.GetDB(ctx, r.db).Save(group); err != nil {
		return "", err
	}
	return group.Id, nil
}

func (r repository) CreateAccountGroup(ctx context.Context, cb func(*models.AccountGroup) error) error {
	collection, err := r.db.FindCachedCollectionByNameOrId("accounts_telegram_groups")
	if err != nil {
		return err
	}
	accountGroup := &models.AccountGroup{}
	accountGroup.SetProxyRecord(core.NewRecord(collection))
	if err := cb(accountGroup); err != nil {
		return err
	}
	return database.GetDB(ctx, r.db).Save(accountGroup)
}
