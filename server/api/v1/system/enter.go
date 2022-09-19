package system

import "akcasbin/server/service"

type SysApiGroup struct {
	DBApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
