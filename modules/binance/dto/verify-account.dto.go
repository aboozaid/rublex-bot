package dto

type VerifyAccountRequestParams struct {
	ApiKey    string
	ApiSecret string
}

type VerifyAccountResponseParams struct {
	AccountID           string
	TelegramBotName     string
	TelegramBotUsername string
}
