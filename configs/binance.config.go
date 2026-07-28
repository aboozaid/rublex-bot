package configs

import (
	"p2ptrader/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type BinanceConfig struct {
	BaseApiUrl            string
	DefaultLimit          string
	RateLimitThreshold    string
	BackoffBaseDuration   string
	MaxConcurrentRequests string
	MaxRequestsPerSecond  string
	RetryCount            string
	RetryMinDuration      string
	RetryMaxDuration      string
	CBMaxFailures         string
	CBCooldownDuration    string
	CBHalfOpenProbeLimit  string
}

func LoadBinanceConfig() BinanceConfig {
	return BinanceConfig{
		BaseApiUrl:            utils.Getenv("BINANCE_BASE_URL", "https://api.binance.com"),
		DefaultLimit:          utils.Getenv("BINANCE_DEFAULT_LIMIT", "6000"),
		RateLimitThreshold:    utils.Getenv("BINANCE_RATE_LIMIT_THRESHOLD", "0.80"),
		BackoffBaseDuration:   utils.Getenv("BINANCE_BACKOFF_BASE_DURATION", "2.5s"),
		MaxConcurrentRequests: utils.Getenv("BINANCE_MAX_CONCURRENT_REQUESTS", "8"),
		MaxRequestsPerSecond:  utils.Getenv("BINANCE_MAX_REQUESTS_PER_SECOND", "5"),
		RetryCount:            utils.Getenv("BINANCE_RETRY_COUNT", "5"),
		RetryMinDuration:      utils.Getenv("BINANCE_RETRY_MIN_DURATION", "1.5s"),
		RetryMaxDuration:      utils.Getenv("BINANCE_RETRY_MAX_DURATION", "30s"),
		CBMaxFailures:         utils.Getenv("BINANCE_CB_MAX_FAILURES", "15"),
		CBCooldownDuration:    utils.Getenv("BINANCE_CB_COOLDOWN_DURATION", "45s"),
		CBHalfOpenProbeLimit:  utils.Getenv("BINANCE_CB_HALF_OPEN_PROBE_LIMIT", "1"),
	}
}

func (c BinanceConfig) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.BaseApiUrl, is.URL),
		validation.Field(&c.RateLimitThreshold, is.Float),
	)
}
