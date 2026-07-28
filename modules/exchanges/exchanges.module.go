package exchanges

import (
	"p2ptrader/common"
	"p2ptrader/modules/binance"
	"p2ptrader/modules/database"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

type module struct {
	service    Service
	controller common.Controller
}

func NewModule(app core.App, binanceModule common.Module[binance.Service]) common.Module[Service] {
	dbService := database.NewModule(app).GetService()
	binanceService := binanceModule.GetService()
	service := newService(app, dbService, binanceService)

	c := newController(service)

	return module{service, c}
}

func (m module) GetService() Service {
	return m.service
}

func (m module) RegisterRoutes(router *router.RouterGroup[*core.RequestEvent]) {
	m.controller.RegisterRoutes(router)
}
