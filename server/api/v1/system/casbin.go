package system

import (
	"akcasbin/forms"
	"akcasbin/models/response"
	"github.com/gin-gonic/gin"
)

type CasbinApi struct{}

func (cas *CasbinApi) GetAllSubjects(c *gin.Context) {
	casbinService.GetAllSubjects()
}

// AddRoleForUserInDomain 赋予用户域角色
func (cas *CasbinApi) AddRoleForUserInDomain(c *gin.Context) {
	var addRoleForm forms.AddRoleForUserInDomain
	if err := c.ShouldBindJSON(&addRoleForm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := casbinService.AddRoleForUserInDomain(&addRoleForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
