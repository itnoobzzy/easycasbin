package core

import (
	"akcasbin/server/global"
	"akcasbin/server/initialize"
	"akcasbin/server/service/system"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CASBIN_CONFIG.System.UseMultipoint || global.CASBIN_CONFIG.System.UseRedis {
		// 初始化redis 服务
		initialize.Redis()
	}
	// 加载jwt
	if global.CASBIN_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CASBIN_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.CASBIN_LOG.Info("server run success on ", zap.String("address", address))
	global.CASBIN_LOG.Error(s.ListenAndServe().Error())
}
