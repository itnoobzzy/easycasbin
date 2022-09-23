package system

import "akcasbin/service"

type SysApiGroup struct {
	DBApi
	UserApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
