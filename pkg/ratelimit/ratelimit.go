package ratelimit

import (
	"errors"
	"sync"
	"time"
)

const (
	defaultMax     = 10
	defaultTimeout = 60
)

type Config struct {
	Max     int
	Timeout int
}

type Hit struct {
	Count     int
	Remaining int
	ResetAt   int
	ResetTime int
}

type Limiter struct {
	Max     int
	timeout int
	hits    map[string]*Hit
	mux     sync.Mutex
}

func New(config ...Config) *Limiter {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Max == 0 {
		cfg.Max = defaultMax
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}

	l := &Limiter{
		Max:     cfg.Max,
		timeout: cfg.Timeout,
		hits:    make(map[string]*Hit),
	}

	return l
}

func (l *Limiter) Hit(key string) (Hit, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	now := int(time.Now().Unix())

	if _, ok := l.hits[key]; !ok {
		l.hits[key] = &Hit{}
	}

	h := l.hits[key]
	h.Count++
	if h.ResetAt == 0 {
		h.ResetAt = now + l.timeout
		time.AfterFunc(time.Duration(l.timeout)*time.Second, func() {
			l.mux.Lock()
			delete(l.hits, key)
			l.mux.Unlock()
		})
	}
	h.ResetTime = h.ResetAt - now
	h.Remaining = l.Max - h.Count
	if h.Remaining < 0 {
		h.Remaining = 0
		return *h, errors.New("hits exceed the max")
	}
	return *h, nil
}
