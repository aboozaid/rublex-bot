package telegramgroups

import (
	"p2ptrader/common"
	"p2ptrader/modules/database"

	"github.com/pocketbase/pocketbase/core"
)

type module struct {
	service Service
}

func NewModule(app core.App) common.Module[Service] {
	dbService := database.NewModule(app).GetService()

	service := newService(app, dbService)
	return module{service}
}

func (m module) GetService() Service {
	return m.service
}
