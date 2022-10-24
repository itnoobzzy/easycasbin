package system

import (
	v1 "akcasbin/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	roleRouter := Router.Group("role")
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("domain_role", roleApi.AddDomainRole)
		roleRouter.PUT("domain_role", roleApi.UpdateDomainRole)
		roleRouter.GET("domain_roles", roleApi.GetAllRoles)
		roleRouter.DELETE("domain_role/:domain_name/:role_name", roleApi.DeleteDomainRole)
	}
	return roleRouter
}
