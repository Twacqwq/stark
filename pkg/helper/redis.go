package helper

import (
	"github.com/FarmerChillax/stark/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(redisConfig *config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		PoolSize:     redisConfig.PoolSize,
		MaxIdleConns: redisConfig.MaxIdle,
	})
	return rdb
}
