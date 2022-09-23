package initialize

import (
	"akcasbin/global"
	"akcasbin/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.CASBIN_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.CasbinRule{},
		models.Role{},
		models.User{},
		models.JwtBlackList{},
	)
	if err != nil {
		global.CASBIN_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.CASBIN_LOG.Info("register table success")
}
