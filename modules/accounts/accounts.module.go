package accounts

import (
	"p2ptrader/common"
	"p2ptrader/modules/database"
	"p2ptrader/modules/kms"

	"github.com/pocketbase/pocketbase/core"
)

type module struct {
	service Service
}

func NewModule(app core.App) common.Module[Service] {
	dbService := database.NewModule(app).GetService()
	kmsService := kms.NewModule().GetService()

	service := newService(app, dbService, kmsService)
	return module{service}
}

func (m module) GetService() Service {
	return m.service
}
