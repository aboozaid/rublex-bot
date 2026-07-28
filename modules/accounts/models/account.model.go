package models

import (
	"fmt"
	"p2ptrader/modules/exchanges/models"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Account struct {
	core.BaseRecordProxy
	/*ID          string
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
	Updated     types.DateTime*/
}

var _ core.RecordProxy = (*Account)(nil)

func (a *Account) Name() string {
	return a.GetString("name")
}

func (a *Account) SetName(name string) {
	a.Set("name", name)
}

func (a *Account) SetNickname(nickname string) {
	a.Set("nickname", nickname)
}

func (a *Account) SetIsActive(isActive bool) {
	a.Set("is_active", isActive)
}

func (a *Account) SetCountryCode(countryCode string) {
	a.Set("country_code", countryCode)
}

func (a *Account) SetDeletedAt(deletedAt time.Time) {
	deletedAtDateTime, _ := types.ParseDateTime(deletedAt)
	a.Set("delete_at", deletedAtDateTime)
}

func (a *Account) SetApiKey(apiKey string) {
	a.Set("api_key", apiKey)
}

func (a *Account) SetApiKeyHash(apiKeyHash string) {
	a.Set("api_key_hashed", apiKeyHash)
}

func (a *Account) SetApiSecret(apiSecret string) {
	a.Set("api_secret", apiSecret)
}

func (a *Account) ApiKey() string {
	return a.GetString("api_key")
}

func (a *Account) ApiSecret() string {
	return a.GetString("api_secret")
}

func (a *Account) Nickname() string {
	return a.GetString("nickname")
}

func (a *Account) UserID() string {
	return a.GetString("user")
}

func (a *Account) Exchange() *models.Exchange {
	var proxy *models.Exchange
	if rel := a.ExpandedOne("exchange"); rel != nil {
		fmt.Println("resolved the relation")
		proxy = &models.Exchange{}
		proxy.SetProxyRecord(rel)
	}
	return proxy
}

func (a *Account) SetExchange(exchangeID string) {
	a.Set("exchange", exchangeID)
}

func (a *Account) SetUser(userID string) {
	a.Set("user", userID)
}

func (a *Account) SetMetadata(metadata string) {
	a.Set("metadata", metadata)
}
