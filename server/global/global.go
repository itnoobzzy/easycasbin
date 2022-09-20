package global

import (
	"akcasbin/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CASBIN_DB     *gorm.DB
	CASBIN_DBLIST map[string]*gorm.DB
	CASBIN_REDIS  *redis.Client
	CASBIN_CONFIG config.Server
	CASBIN_VP     *viper.Viper
	// CASBIN_LOG
	CASBIN_LOG *zap.Logger

	BlackCache local_cache.Cache
)
