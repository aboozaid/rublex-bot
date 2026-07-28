package exchanges

import (
	"context"
	"p2ptrader/modules/database"
	"p2ptrader/modules/exchanges/models"

	"github.com/pocketbase/pocketbase/core"
)

var collectionKey = "exchanges"

type repo interface {
	GetExchangeByID(ctx context.Context, exchangeID string) (*models.Exchange, error)
}

type repository struct {
	db core.App
}

func newRepository(db core.App) repo {
	return repository{db}
}

func (r repository) GetExchangeByID(ctx context.Context, exchangeID string) (*models.Exchange, error) {
	exchange := &models.Exchange{}
	record, err := database.GetDB(ctx, r.db).FindRecordById(collectionKey, exchangeID)
	if err != nil {
		return nil, err
	}
	exchange.Record = record
	return exchange, nil
}
