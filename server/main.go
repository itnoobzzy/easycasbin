package main

import (
	"akcasbin/core"
	"akcasbin/global"
	"akcasbin/initialize"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	global.CASBIN_VP = core.Viper()
	global.CASBIN_LOG = core.Zap()
	zap.ReplaceGlobals(global.CASBIN_LOG)
	global.CASBIN_DB = initialize.Gorm()
	initialize.DBList()
	if global.CASBIN_DB != nil {
		initialize.RegisterTables(global.CASBIN_DB)
		db, _ := global.CASBIN_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
