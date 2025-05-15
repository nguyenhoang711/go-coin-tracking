package initialize

import (
	"context"
	"fmt"

	"github.com/nguyenhoang711/go-coin-tracking/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	_ctx = context.Background()
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(_ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: %v", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Connecting redis successfully!")
	global.Rdb = rdb
}
