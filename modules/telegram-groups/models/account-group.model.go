package models

import "github.com/pocketbase/pocketbase/core"

type AccountGroup struct {
	core.BaseRecordProxy
}

var _ core.RecordProxy = (*AccountGroup)(nil)

func (a *AccountGroup) GroupID() string {
	return a.GetString("group")
}

func (a *AccountGroup) AccountID() string {
	return a.GetString("account")
}

func (a *AccountGroup) SetGroup(groupID string) {
	a.Set("group", groupID)
}

func (a *AccountGroup) SetAccount(accountID string) {
	a.Set("account", accountID)
}
