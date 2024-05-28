package db

import (
	"MyBalance/internal/core/db/db_in_memory"
	"MyBalance/internal/core/db/interface_mono_db"
	"MyBalance/internal/core/db/internal_redis"
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/logger"
	"MyBalance/internal/projkeys"
	"github.com/redis/go-redis/v9"
)

var db interface_mono_db.DBMono

func Init(ctx context.Context) error {
	dbType, _ := ctx.GetString(projkeys.DatabaseType)
	if dbType == "" {
		logger.PrintInfo(ctx, "Using db_in_memory as database")

		db = db_in_memory.NewDatabaseMonoInMemory()
	} else if dbType == "redis" {
		logger.PrintInfo(ctx, "Using redis as database")

		url, err := ctx.GetString(projkeys.RedisCloudUrl)
		if err != nil {
			logger.PrintError(ctx, "env missing "+projkeys.RedisCloudUrl, err)
			return err
		}

		opts, err := redis.ParseURL(url)
		if err != nil {
			logger.PrintError(ctx, "redis parseUrl error ", err)
			return err
		}

		db = internal_redis.NewClient(ctx, opts)
	}

	return nil
}
