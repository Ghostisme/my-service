package global

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)
