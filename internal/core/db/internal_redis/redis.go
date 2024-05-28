package internal_redis

import (
	"MyBalance/internal/core/db/interface_mono_db"
	"MyBalance/internal/http/context"
	"github.com/redis/go-redis/v9"
)

var _ interface_mono_db.DBMono = (*redisStruct)(nil)

type redisStruct struct {
	db *redis.Client
}

const (
	CardToUserTable int = 0
)

func (r *redisStruct) GetCard(ctx context.Context, key string) (string, error) {
	val, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisStruct) SetCard(ctx context.Context, key string, value string) error {
	return r.db.Set(ctx, key, value, 0).Err()
}

func (r *redisStruct) GetUsers(ctx context.Context) []string {
	keys := r.db.Keys(ctx, "*")

	return keys.Val()
}

func NewClient(ctx context.Context, opt *redis.Options) interface_mono_db.DBMono {
	return &redisStruct{
		db: redis.NewClient(opt),
	}
}
