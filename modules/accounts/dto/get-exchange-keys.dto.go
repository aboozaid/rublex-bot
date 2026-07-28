package dto

import "p2ptrader/modules/exchanges/dto"

type GetEncryptedExchangeKeysResponseParams struct {
	Exchange  dto.Exchange `json:"exchange"`
	ApiKey    string       `json:"api_key"`
	ApiSecret string       `json:"api_secret"`
}
