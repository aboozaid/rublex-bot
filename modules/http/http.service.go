package http

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

type rateLimiter struct {
	mu        sync.Mutex
	used      int
	limit     int
	semaphore chan struct{}
	rpsTokens chan struct{}
	rpsTicker *time.Ticker
	cfg       HttpConfig
	strategy  HttpStrategy
}

type cbState int

const (
	stateClosed = iota
	stateOpened
	stateHalfOpen
)

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open: requests are currently blocked")

type circuitBreaker struct {
	mu       sync.Mutex
	state    cbState
	Cfg      CircuitBreakerConfig
	failures int
	probes   int
	openTime time.Time
}

type Service struct {
	Client *req.Client
}

func newService(client *req.Client, strategy HttpStrategy, cfg HttpConfig) Service {
	sem := make(chan struct{}, cfg.MaxConcurrentRequests)
	for i := 0; i < cfg.MaxConcurrentRequests; i++ {
		sem <- struct{}{}
	}

	var tokens chan struct{}
	var ticker *time.Ticker
	if cfg.MaxRequestsPerSecond > 0 {
		tokens = make(chan struct{}, cfg.MaxRequestsPerSecond)
		for i := 0; i < cfg.MaxRequestsPerSecond; i++ {
			tokens <- struct{}{}
		}
		ticker = time.NewTicker(time.Second)
		go func() {
			for range ticker.C {
				for i := 0; i < cfg.MaxRequestsPerSecond; i++ {
					select {
					case tokens <- struct{}{}:
					default:
						// pool is full stop refilling
					}
				}
			}
		}()
	}

	rl := &rateLimiter{limit: cfg.DefaultLimit, cfg: cfg, strategy: strategy, semaphore: sem, rpsTokens: tokens, rpsTicker: ticker}
	cb := &circuitBreaker{Cfg: cfg.CircuitBreaker}
	client.WrapRoundTripFunc(func(rt req.RoundTripper) req.RoundTripFunc {
		return func(req *req.Request) (resp *req.Response, err error) {
			// CURRENTLY: if we hit circuit breaker opened gate requests will retry silently after that dropped entirly
			// REVIEW: We could improve that by a job queue to re-schedule after retrying attemps fail
			if cbErr := cb.AllowRequest(); cbErr != nil {
				return nil, cbErr
			}
			if rlErr := rl.RequestMiddleware(client, req); rlErr != nil {
				return nil, rlErr
			}
			resp, err = rt.RoundTrip(req)

			//isSuccess := err == nil && resp.GetStatusCode() < 500 && (resp.GetStatusCode() != 429 || resp.GetStatusCode() != 418)
			isSuccess := strategy.IsRequestSuccess(resp, err)
			cb.RecordResult(isSuccess)

			if err == nil {
				rl.ResponseMiddleware(client, resp)
			}
			return resp, err
		}
	})
	return Service{client}
}

func (rl *rateLimiter) RequestMiddleware(client *req.Client, r *req.Request) error {
	if rl.rpsTokens != nil {
		select {
		case <-r.Context().Done():
			return r.Context().Err()
		case <-rl.rpsTokens:
			// request token acquired otherwise wait for a token
		}
	}

	select {
	case <-r.Context().Done():
		return r.Context().Err()
	case <-rl.semaphore:
		// request executed concurrently otherwise wait for a slot
	}

	if err := rl.strategy.OnBeforeRequest(r); err != nil {
		return err
	}

	r.OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		rl.semaphore <- struct{}{}
		return nil
	})

	rl.mu.Lock()
	used := rl.used
	limit := rl.limit
	rl.mu.Unlock()

	if limit == 0 {
		return nil
	}

	usageRatio := float64(used) / float64(limit)

	if usageRatio >= rl.cfg.RateLimitThreshold {
		exceesRatio := usageRatio - (rl.cfg.RateLimitThreshold - 0.05)
		sleepTime := time.Duration(float64(rl.cfg.BackoffBaseDuration) * exceesRatio)

		log := fmt.Sprintf("[RateLimiter] Target Threshold Warning (%.2f%% >= %.2f%%). Queue backing off for %v\n",
			usageRatio*100, rl.cfg.RateLimitThreshold*100, sleepTime)

		slog.Info(log)
		select {
		case <-r.Context().Done():
			return r.Context().Err()
		case <-time.After(sleepTime):
		}
	}

	return nil
}

func (rl *rateLimiter) ResponseMiddleware(client *req.Client, resp *req.Response) error {
	used, limit := rl.strategy.ParseHeaders(resp)

	if limit > 0 && used > 0 {
		rl.mu.Lock()
		rl.used = used
		rl.limit = limit
		rl.mu.Unlock()
	}
	return nil
}

func (cb *circuitBreaker) AllowRequest() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case stateClosed:
		return nil
	case stateOpened:
		if time.Since(cb.openTime) >= cb.Cfg.CooldownDuration {
			cb.state = stateHalfOpen
			cb.probes = 0
			return nil
		}
		return ErrCircuitBreakerOpen
	case stateHalfOpen:
		if cb.probes < cb.Cfg.HalfOpenProbeLimit {
			cb.probes++
			return nil
		}
		return ErrCircuitBreakerOpen
	}
	return nil
}

func (cb *circuitBreaker) RecordResult(success bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if success {
		cb.failures = 0
		cb.state = stateClosed
		return
	}

	cb.failures++
	if cb.state == stateClosed && cb.failures >= cb.Cfg.MaxFailures {
		cb.state = stateOpened
		cb.openTime = time.Now()
	} else if cb.state == stateHalfOpen {
		cb.state = stateOpened
		cb.openTime = time.Now()
	}
}
