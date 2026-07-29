package dto

type CreateUserParams struct {
	TelegramID	string
	TelegramUsername	*string
	FirstName	string
	LastName	string
	CountryCode	string
}
