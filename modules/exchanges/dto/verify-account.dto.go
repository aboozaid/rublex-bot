package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type VerifyAccountRequestParams struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func (r VerifyAccountRequestParams) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ApiKey, validation.Required),
		validation.Field(&r.ApiSecret, validation.Required),
	)
}

type VerifyAccountResponseParams struct {
	AccountID           string `json:"account_id"`
	TelegramBotName     string `json:"telegram_bot_name"`
	TelegramBotUsername string `json:"telegram_bot_username"`
}
