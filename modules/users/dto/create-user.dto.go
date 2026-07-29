package dto

type CreateUserRequestParams struct {
	TelegramID	string
	TelegramUsername	*string
	FirstName	string
	LastName	string
	LanguageCode	string
}