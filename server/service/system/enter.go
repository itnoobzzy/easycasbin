package system

type ServiceGroup struct {
	InitDBService
	UserService
	JwtService
	CasbinService
	RoleService
}
