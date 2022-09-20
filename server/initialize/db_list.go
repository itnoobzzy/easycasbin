package initialize

import (
	"akcasbin/config"
	"akcasbin/global"
	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.CASBIN_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{
				GeneralDB: info.GeneralDB,
			})
		case "pgsql":
			dbMap[info.AliasName] = GormPgsqlByConfig(config.Pgsql{
				GeneralDB: info.GeneralDB,
			})
		default:
			continue
		}
	}
	if sysDB, ok := dbMap[sys]; ok {
		global.CASBIN_DB = sysDB
	}
	global.CASBIN_DBLIST = dbMap
}
