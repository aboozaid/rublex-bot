package binance

import (
	"log/slog"
	"os"
	"p2ptrader/common"
	"p2ptrader/modules/accounts"
	"p2ptrader/modules/cache"
	"p2ptrader/modules/config"
	"p2ptrader/modules/http"

	"github.com/pocketbase/pocketbase/core"
)

type module struct {
	service Service
}

/*
Binance rate limits

  - Every request will contain X-MBX-USED-WEIGHT-(intervalNum)(intervalLetter) in the response headers which has the current used weight for the IP for all request rate limiters defined.

  - Each route has a weight which determines for the number of requests each endpoint counts for. Heavier endpoints and endpoints that do operations on multiple symbols will have a heavier weight

  - When a 429 is received, it's your obligation as an API to back off and not spam the API.

  - Repeatedly violating rate limits and/or failing to back off after receiving 429s will result in an automated IP ban (HTTP status 418).

  - adv	All /sapi/v1/c2c/ endpoints 12000 RPM

  - sapi	Signed SAPI requests (general) 12000 RPM

  - spot	/api/v3/ endpoints (e.g. server time) 12000 RPM
*/
func NewModule(app core.App, cacheModule common.Module[cache.Service], accountsModule common.Module[accounts.Service]) common.Module[Service] {
	configService := config.NewModule().GetService()

	strategy := newStrategy()
	binancecfg := configService.Binance()
	parser := &config.EnvParser{}
	httpConfig := http.HttpConfig{
		DefaultBaseURL:        binancecfg.BaseApiUrl,
		DefaultLimit:          parser.Int(binancecfg.DefaultLimit, "BINANCE_DEFAULT_LIMIT"),
		RateLimitThreshold:    parser.Float(binancecfg.RateLimitThreshold, "BINANCE_RATE_LIMIT_THRESHOLD"),
		BackoffBaseDuration:   parser.Duration(binancecfg.BackoffBaseDuration, "BINANCE_BACKOFF_BASE_DURATION"),
		MaxConcurrentRequests: parser.Int(binancecfg.MaxConcurrentRequests, "BINANCE_MAX_CONCURRENT_REQUESTS"),
		MaxRequestsPerSecond:  parser.Int(binancecfg.MaxRequestsPerSecond, "BINANCE_MAX_REQUESTS_PER_SECOND"),

		RetryCount:       parser.Int(binancecfg.RetryCount, "BINANCE_RETRY_COUNT"),
		RetryMinDuration: parser.Duration(binancecfg.RetryMinDuration, "BINANCE_RETRY_MIN_DURATION"),
		RetryMaxDuration: parser.Duration(binancecfg.RetryMaxDuration, "BINANCE_RETRY_MAX_DURATION"),
		EnableDevMode:    configService.App().Environment != "production",

		CircuitBreaker: http.CircuitBreakerConfig{
			MaxFailures:        parser.Int(binancecfg.CBMaxFailures, "BINANCE_CB_MAX_FAILURES"),
			CooldownDuration:   parser.Duration(binancecfg.CBCooldownDuration, "BINANCE_CB_COOLDOWN_DURATION"),
			HalfOpenProbeLimit: parser.Int(binancecfg.CBHalfOpenProbeLimit, "BINANCE_CB_HALF_OPEN_PROBE_LIMIT"),
		},
	}
	if parser.Err != nil {
		slog.Error("unable to parse binance configs", "error", parser.Err)
		os.Exit(1)
	}
	httpService := http.NewModule(strategy, httpConfig).GetService()
	cacheService := cacheModule.GetService()
	accountsService := accountsModule.GetService()
	service := newService(app, cacheService, httpService, accountsService)
	return module{service}
}

func (m module) GetService() Service {
	return m.service
}
