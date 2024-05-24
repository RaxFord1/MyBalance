package limiter

import (
	"MyBalance/internal/http/requesto"
	"golang.org/x/time/rate"
	"sync"
)

type EventLimiter struct {
	limiters map[string]*rate.Limiter //
	mu       sync.Mutex
	rate     rate.Limit
	burst    int
}

// NewEventLimiter creates a new EventLimiter.
func NewEventLimiter(r rate.Limit, b int) *EventLimiter {
	return &EventLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    b,
	}
}

func (l *EventLimiter) GetLimiter(userID string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	limiter, exists := l.limiters[userID]
	if !exists {
		limiter = rate.NewLimiter(l.rate, l.burst)
		l.limiters[userID] = limiter
	}
	return limiter
}

func (l *EventLimiter) Allow(key string) error {
	limiter := l.GetLimiter(key)
	if !limiter.Allow() {
		return requesto.TooManyRequests
	}
	return nil
}
