package models

/*type Exchange struct {
	ID       string
	Logo     string
	Website  string
	Name     string
	IsActive bool
	Created  types.DateTime
	Updated  types.DateTime
}*/

import (
	"github.com/pocketbase/pocketbase/core"
)

type Exchange struct {
	core.BaseRecordProxy
}

var _ core.RecordProxy = (*Exchange)(nil)

type exchangeType string

const (
	Binance exchangeType = "binance"
	Bybit   exchangeType = "bybit"
	Okx     exchangeType = "okx"
)

func (a *Exchange) Name() string {
	return a.GetString("name")
}

func (a *Exchange) SetName(name string) {
	a.Set("name", name)
}

func (a *Exchange) SetLogo(logo string) {
}

func (a *Exchange) SetIsActive(isActive bool) {
	a.Set("is_active", isActive)
}

func (a *Exchange) SetWebsite(website string) {
	a.Set("website", website)
}

func (a *Exchange) ExchangeType() exchangeType {
	switch a.Name() {
	case "Binance":
		return Binance
	case "ByBit":
		return Bybit
	default:
		return Okx
	}
}
