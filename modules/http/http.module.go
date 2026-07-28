package http

import (
	"p2ptrader/common"

	"github.com/imroc/req/v3"
)

type module struct {
	service Service
}

func NewModule(strategy HttpStrategy, cfg HttpConfig) common.Module[Service] {
	client := req.C().
		SetCommonRetryCount(cfg.RetryCount).
		SetCommonRetryBackoffInterval(cfg.RetryMinDuration, cfg.RetryMaxDuration).
		SetCommonRetryCondition(func(resp *req.Response, err error) bool {
			return strategy.ShouldRetry(resp, err)
		}).
		SetBaseURL(cfg.DefaultBaseURL)

	if cfg.EnableDevMode {
		client.DevMode()
	}

	service := newService(client, strategy, cfg)

	return module{service}
}

func (m module) GetService() Service {
	return m.service
}
