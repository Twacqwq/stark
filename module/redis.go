package module

import (
	"context"

	"github.com/FarmerChillax/stark"
	"github.com/redis/go-redis/v9"
)

type redisConn struct {
	client *redis.Client
}

func (rc *redisConn) Get(ctx context.Context) *redis.Client {
	return rc.client
}

func RegisterRedis(rdb *redis.Client) {
	stark.Redis = &redisConn{client: rdb}
}
