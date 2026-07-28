package http

import (
	"time"

	"github.com/imroc/req/v3"
)

type CircuitBreakerConfig struct {
	MaxFailures        int
	CooldownDuration   time.Duration
	HalfOpenProbeLimit int
}

type HttpConfig struct {
	DefaultBaseURL        string
	DefaultLimit          int
	RateLimitThreshold    float64
	BackoffBaseDuration   time.Duration
	MaxConcurrentRequests int
	MaxRequestsPerSecond  int
	RetryCount            int
	RetryMinDuration      time.Duration
	RetryMaxDuration      time.Duration

	EnableDevMode bool

	CircuitBreaker CircuitBreakerConfig
}

type HttpStrategy interface {
	ParseHeaders(resp *req.Response) (int, int)
	ShouldRetry(resp *req.Response, err error) bool
	OnBeforeRequest(req *req.Request) error
	IsRequestSuccess(resp *req.Response, err error) bool
}
