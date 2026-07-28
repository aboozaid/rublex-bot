package database

import (
	"p2ptrader/common"
	"sync"

	"github.com/pocketbase/pocketbase/core"
)

type module struct {
	service Service
}

var (
	instance *module
	once     sync.Once
)

func NewModule(app core.App) common.Module[Service] {
	once.Do(func() {
		instance = &module{NewService(app)}
	})
	return instance
}

func (m *module) GetService() Service {
	return m.service
}
