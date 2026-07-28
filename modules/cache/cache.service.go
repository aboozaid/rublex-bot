package cache

import (
	"context"
	"sync"
	"time"
)

type Service interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Remove(key string)

	start() error
	stop() error
}

type entry struct {
	data      string
	expiresAt time.Time
}

type service struct {
	entries sync.Map
	ttl     time.Duration
	cancel  context.CancelFunc
	stopped chan struct{}
}

func newService(ttl time.Duration) Service {
	return &service{ttl: ttl, stopped: make(chan struct{})}
}

func (s *service) Set(key, value string) {
	s.entries.Store(key, entry{value, time.Now().Add(s.ttl)})
}
func (s *service) Get(key string) (string, bool) {
	v, ok := s.entries.Load(key)
	if !ok {
		return "", false
	}
	e := v.(entry)
	if time.Now().After(e.expiresAt) {
		s.entries.Delete(key)
		return "", false
	}
	return e.data, true
}
func (s *service) Remove(key string) {
	s.entries.Delete(key)
}

func (s *service) cleanupLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer close(s.stopped)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			s.entries.Range(func(key, value any) bool {
				if now.After(value.(entry).expiresAt) {
					s.entries.Delete(key)
				}
				return true
			})
		}
	}
}

func (s *service) start() error {
	runCtx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	go s.cleanupLoop(runCtx, s.ttl/2)
	return nil
}

func (s *service) stop() error {
	if s.cancel != nil {
		s.cancel()
	}

	<-s.stopped
	return nil
}
