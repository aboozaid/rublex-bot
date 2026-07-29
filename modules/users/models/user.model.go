package models

import (
	"github.com/pocketbase/pocketbase/core"
)

// type User struct {
// 	ID          string
// 	Exchange    models.Exchange
// 	ApiKey      string
// 	ApiSecret   string
// 	Name        *string
// 	Nickname    *string
// 	CountryCode *string
// 	TgGroupID   *string
// 	IsActive    bool
// 	Metadata    any
// 	DeletedAt   *types.DateTime
// 	Created     types.DateTime
// 	Updated     types.DateTime
// }
type User struct {
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

var _ core.RecordProxy = (*User)(nil)

func (u *User) Name() string {
	return u.GetString("name")
}

func (u *User) SetName(name string) {
	u.Set("name", name)
}

func (u *User) Firstname() string {
	return u.GetString("first_name")
}

func (u *User) SetFirstName(firstname string) {
	u.Set("first_name", firstname)
}

func (u *User) Lastname() string {
	return u.GetString("last_name")
}

func (u *User) SetLastName(lastname string) {
	u.Set("last_name", lastname)
}

func (u *User) TelegramUsername() string {
	return u.GetString("telegram_username")
}

func (u *User) SetTelegramUsername(tgUsername string) {
	u.Set("telegram_username", tgUsername)
}
func (u *User) TelegramID() string {
	return u.GetString("telegram_id")
}
func (u *User) SetTelegramID(tgID string) {
	u.Set("telegram_id", tgID)
}
func (u *User) LanguageCode() string {
	return u.GetString("language_code")
}
func (u *User) SetLanguageCode(languageCode string) {
	u.Set("language_code", languageCode)
}
func (u *User) PhotoUrl() string {
	return u.GetString("photo_url")
}
func (u *User) SetPhotoUrl(photoUrl string) {
	u.Set("photo_url", photoUrl)
}
