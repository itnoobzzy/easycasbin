package system

import (
	"akcasbin/forms"
	"akcasbin/models/response"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func (ra *RoleApi) AddDomainRole(c *gin.Context) {
	var addDomainRoleForm forms.AddDomainRole
	if err := c.ShouldBindJSON(&addDomainRoleForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}

	role, err := roleService.AddRole(&addDomainRoleForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(role, "添加成功!", c)
}
