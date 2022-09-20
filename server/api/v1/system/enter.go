package system

import "akcasbin/service"

type SysApiGroup struct {
	DBApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
