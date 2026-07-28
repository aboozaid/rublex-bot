package users

import (
	"p2ptrader/common"

	"github.com/pocketbase/pocketbase/core"
)

type module struct {
	service Service
}

func NewModule(app core.App) common.Module[Service] {
	service := NewService(app)
	return module{service}
}

func (m module) GetService() Service {
	return m.service
}
