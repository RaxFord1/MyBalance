package mono_balance

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/limiter"
	"golang.org/x/time/rate"
	"time"
)

var (
	limiterAll    *limiter.EventLimiter
	limiterClient *limiter.EventLimiter
)

func Init(ctx context.Context) error {
	limiterAll = limiter.NewEventLimiter(rate.Every(time.Minute), 5)

	limiterClient = limiter.NewEventLimiter(rate.Every(time.Minute), 1)

	return nil
}
