package initialize

import (
	"easycasbin/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CASBIN_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.CASBIN_LOG.Error("redis conn ping failed, err:", zap.Error(err))
	} else {
		global.CASBIN_LOG.Info("redis conn ping response:", zap.String("pong", pong))
		global.CASBIN_REDIS = client
	}
}
