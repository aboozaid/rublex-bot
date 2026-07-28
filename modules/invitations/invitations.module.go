package invitations

import (
	"p2ptrader/common"
	"p2ptrader/modules/config"
	"p2ptrader/modules/database"
	"p2ptrader/modules/users"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type module struct {
	service    Service
	controller common.Controller
	handler    common.Handler
}

func NewModule(app core.App, usersModule common.Module[users.Service]) common.Module[Service] {
	dbService := database.NewModule(app).GetService()
	service := newService(app, usersModule.GetService(), dbService)

	c := newController(service)
	h := newHandler(service, config.NewModule().GetService())

	return module{service, c, h}
}

func (m module) GetService() Service {
	return m.service
}

func (m module) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
	m.controller.RegisterRoutes(router)
}

func (m module) RegisterHandlers(client *telegram.Client) {
	m.handler.RegisterHandlers(client)
}
