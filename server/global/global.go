package global

import (
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/cache/singleflight"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"akcasbin/config"
)

var (
	// system global
	CASBIN_DB                  *gorm.DB
	CASBIN_DBLIST              map[string]*gorm.DB
	CASBIN_REDIS               *redis.Client
	CASBIN_CONFIG              config.Server
	CASBIN_VP                  *viper.Viper
	CASBIN_LOG                 *zap.Logger
	CASBIN_CONCURRENCY_CONTROL = &singleflight.Group{}
	FORM_TRANS                 ut.Translator

	BlackCache local_cache.Cache

	// casbin enforcer
	CASBIN_ENFORCER *casbin.CachedEnforcer
)
