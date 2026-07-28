package cache

import (
	"log/slog"
	"os"
	"p2ptrader/common"
	"p2ptrader/modules/config"
	"sync"
)

var (
	instance *module
	once     sync.Once
)

type module struct {
	service Service
}

func NewModule() common.Module[Service] {
	once.Do(func() {
		configService := config.NewModule().GetService()
		p := &config.EnvParser{}
		cacheTTL := p.Duration(configService.App().TTlMemoryCacheDuration, "MAX_TTL_MEMORY_CACHE_DURATION")
		if p.Err != nil {
			slog.Error("unable to parse cache config", "error", p.Err)
			os.Exit(1)
		}
		instance = &module{service: newService(cacheTTL)}
	})
	return instance
}

func (m *module) GetService() Service {
	return m.service
}

func (m *module) OnAppStart() error {
	return m.service.start()
}

func (m *module) OnAppStop() error {
	return m.service.stop()
}
