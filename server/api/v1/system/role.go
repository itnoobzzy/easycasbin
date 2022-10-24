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

// UpdateDomainRole 更新域角色信息
func (ra *RoleApi) UpdateDomainRole(c *gin.Context) {
	var form forms.UpdateDomainRole
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	role, err := roleService.UpdateRoleInfo(&form)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(role, c)
}

// DeleteDomainRole 删除对应域的角色
func (ra *RoleApi) DeleteDomainRole(c *gin.Context) {
	var DeleteDomainRole forms.DeleteDomainRole
	if err := c.ShouldBindUri(&DeleteDomainRole); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := roleService.DeleteRole(&DeleteDomainRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
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
