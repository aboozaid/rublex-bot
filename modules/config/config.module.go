package config

import (
	"p2ptrader/common"
	"sync"
)

type module struct {
	service Service
}

var (
	instance *module
	once     sync.Once
)

func NewModule() common.Module[Service] {
	once.Do(func() {
		instance = &module{newService()}
	})
	return instance
}

func (m *module) GetService() Service {
	return m.service
}
