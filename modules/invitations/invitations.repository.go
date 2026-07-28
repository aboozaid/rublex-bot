package invitations

import (
	"context"
	"p2ptrader/modules/database"

	"github.com/pocketbase/pocketbase/core"
)

type repo interface {
	FindInvitationByCode(ctx context.Context, code string) (*core.Record, error)
	UpdateInvitation(ctx context.Context, cb func() (*core.Record, error)) error
	CreateInvitation(ctx context.Context, cb func(*core.Record) error) error
}

type repository struct {
	db core.App
}

const COLLECTION_NAME = "invitations"

func NewRepository(db core.App) repo {
	return repository{db}
}

func (r repository) FindInvitationByCode(ctx context.Context, code string) (*core.Record, error) {
	invitation, err := database.GetDB(ctx, r.db).FindFirstRecordByData(COLLECTION_NAME, "code", code)
	return invitation, err
}

func (r repository) UpdateInvitation(ctx context.Context, cb func() (*core.Record, error)) error {
	invitation, err := cb()
	if err != nil {
		return err
	}
	return database.GetDB(ctx, r.db).Save(invitation)
}

func (r repository) CreateInvitation(ctx context.Context, cb func(*core.Record) error) error {
	collection, err := r.db.FindCollectionByNameOrId(COLLECTION_NAME)
	if err != nil {
		return err
	}
	record := core.NewRecord(collection)
	if err := cb(record); err != nil {
		return err
	}
	return database.GetDB(ctx, r.db).Save(record)
}
