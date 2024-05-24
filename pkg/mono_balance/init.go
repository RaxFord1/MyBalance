package mono_balance

import (
	"MyBalance/internal/context"
	"MyBalance/internal/limiter"
)

var (
	limiterAll    *limiter.EventLimiter
	limiterClient *limiter.EventLimiter
)

func Init(ctx context.Context) error {
	limiterAll = limiter.NewEventLimiter(1, 5)

	limiterClient = limiter.NewEventLimiter(1, 1)

	return nil
}
