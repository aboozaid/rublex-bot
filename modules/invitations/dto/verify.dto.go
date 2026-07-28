package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type VerifyRequestParams struct {
	Code       string `json:"code"`
	TelegramID string `json:"telegram_id"`
}

func (r VerifyRequestParams) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Code, validation.Required),
		validation.Field(&r.TelegramID, validation.Required),
	)
}

type VerifyResponseParams struct {
	Verified bool `json:"verified"`
}
