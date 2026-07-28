package models

import (
	"p2ptrader/modules/exchanges/models"

	"github.com/pocketbase/pocketbase/tools/types"
)

type User struct {
	ID          string
	Exchange    models.Exchange
	ApiKey      string
	ApiSecret   string
	Name        *string
	Nickname    *string
	CountryCode *string
	TgGroupID   *string
	IsActive    bool
	Metadata    any
	DeletedAt   *types.DateTime
	Created     types.DateTime
	Updated     types.DateTime
}
