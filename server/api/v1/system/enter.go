package system

import "akcasbin/service"

type SysApiGroup struct {
	DBApi
	UserApi
	CasbinApi
	RoleApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	roleService   = service.ServiceGroupApp.SystemServiceGroup.RoleService
)
