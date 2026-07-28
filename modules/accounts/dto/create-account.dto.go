package dto

type CreateAccountRequestParams struct {
	ExchangeID  string
	ApiKey      string
	ApiSecret   string
	Name        string
	Nickname    *string
	CountryCode *string
	//Metadata    map[string]any
}
