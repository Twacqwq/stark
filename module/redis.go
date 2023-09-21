package module

import (
	"context"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type redisConn struct {
	client *redis.Client
}

func (rc *redisConn) Get(ctx context.Context) *redis.Client {
	return rc.client
}

func RegisterRedis(app *stark.Application) error {
	redisConf := app.Config.Redis
	rdb := utils.NewRedis(redisConf)
	stark.Redis = &redisConn{client: rdb}
	return nil
}
