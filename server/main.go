package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"akcasbin/core"
	"akcasbin/global"
	"akcasbin/initialize"
)

func main() {
	global.CASBIN_VP = core.Viper()
	global.CASBIN_LOG = core.Zap()
	zap.ReplaceGlobals(global.CASBIN_LOG)
	global.CASBIN_DB = initialize.Gorm()
	global.CASBIN_ENFORCER = core.Enforcer()
	initialize.DBList()
	// 初始化表单校验翻译器
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	if global.CASBIN_DB != nil {
		initialize.RegisterTables(global.CASBIN_DB)
		db, _ := global.CASBIN_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
