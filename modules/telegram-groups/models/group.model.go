package models

import (
	"github.com/pocketbase/pocketbase/core"
)

type Group struct {
	core.BaseRecordProxy
}

var _ core.RecordProxy = (*Group)(nil)

func (a *Group) GroupID() string {
	return a.GetString("group_id")
}

func (a *Group) SetGroupID(groupID string) {
	a.Set("group_id", groupID)
}

func (a *Group) SetTitle(title string) {
	a.Set("title", title)
}

func (a *Group) SetUsername(username string) {
	a.Set("username", username)
}

func (a *Group) Title() string {
	return a.GetString("title")
}

func (a *Group) Username() string {
	return a.GetString("username")
}

func (a *Group) SetUser(userID string) {
	a.Set("user", userID)
}
