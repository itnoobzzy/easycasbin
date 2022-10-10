package system

import (
	"akcasbin/forms"
	"akcasbin/models/response"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

// AddDomainRole 添加对应域的角色
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

// GetAllRoles 查询域下所有角色
func (ra *RoleApi) GetAllRoles(c *gin.Context) {
	var getDomainRolesForm forms.GetAllRoles
	if err := c.ShouldBind(&getDomainRolesForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	roles := roleService.GetDomainRoles(&getDomainRolesForm)
	response.OkWithDetailed(roles, "ok", c)
}
