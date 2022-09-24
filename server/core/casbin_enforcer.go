package core

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"

	"akcasbin/global"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

// Enforcer 获取casbin 的enforcer
func Enforcer() (enforcer *casbin.CachedEnforcer) {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.CASBIN_DB)
		m, err := model.NewModelFromFile("./server/casbin_rbac_domain.conf")
		if err != nil {
			zap.L().Error("casbin model load failed!", zap.Error(err))
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
