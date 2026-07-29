package telegram

import (
	"p2ptrader/common"
	"p2ptrader/modules/accounts"
	telegramgroups "p2ptrader/modules/telegram-groups"
	"p2ptrader/modules/users"

	"github.com/amarnathcjd/gogram/telegram"
)

type module struct {
	service Service
	handler common.Handler
}

func NewModule(accountsModule common.Module[accounts.Service], telegramGroups common.Module[telegramgroups.Service], usersModule common.Module[users.Service]) common.Module[Service] {
	service := newService(accountsModule.GetService(), telegramGroups.GetService(), usersModule.GetService())

	h := newHandler(service)

	return module{service, h}
}

func (m module) GetService() Service {
	return m.service
}

func (m module) RegisterHandlers(client *telegram.Client) {
	m.handler.RegisterHandlers(client)
}
