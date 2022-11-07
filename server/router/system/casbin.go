package system

import (
	v1 "akcasbin/api/v1"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	casbinRouter := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.GET("all_subjects", casbinApi.GetAllSubjects)
		casbinRouter.GET("all_domains", casbinApi.GetAllDomains)
		casbinRouter.POST("domain", casbinApi.AddDomain)
		casbinRouter.DELETE("domain/:domain_name", casbinApi.DeleteDomain)

		casbinRouter.GET("roles_in_domain", casbinApi.GetRolesForUserInDomain)
		casbinRouter.POST("enforce_policies", casbinApi.BatchEnforce)

		casbinRouter.POST("domain_permissions", casbinApi.AddPermissionsForSubInDomain)
		casbinRouter.PUT("domain_permissions", casbinApi.UpdatePermissionsForSubInDomain)
		casbinRouter.GET("domain_permissions", casbinApi.GetPermissionsForSubInDomain)
		casbinRouter.DELETE("domain_permissions", casbinApi.DeletePermissionsForSubInDomain)
	}
	return casbinRouter
}
