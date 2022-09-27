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
		casbinRouter.POST("add_domain_role", casbinApi.AddRoleForUserInDomain)
		casbinRouter.POST("add_policy", casbinApi.AddPolicy)
	}
	return casbinRouter
}
