package dto

type BinanceResponse[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Total   *int64 `json:"total"`
	Success bool   `json:"success"`
}

type BinanceErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type BinanceUser struct {
	Nickname         string `json:"nickname"`
	KycFullname      string `json:"kycFullName"`
	CountryCode      string `json:"countryCode"`
	IsNicknameExists bool   `json:"existsNickname"`
}
