package system

import (
	"akcasbin/forms"
	"akcasbin/models/response"
	"akcasbin/utils"
	"github.com/gin-gonic/gin"
)

type CasbinApi struct{}

// GetAllSubjects 获取所有鉴权对象
func (cas *CasbinApi) GetAllSubjects(c *gin.Context) {
	res := casbinService.GetAllSubjects()
	response.OkWithData(res, c)
}

// GetAllDomains 获取所有域
func (cas *CasbinApi) GetAllDomains(c *gin.Context) {
	res, _ := casbinService.GetAllDomains()
	response.OkWithData(res, c)
}

// AddDomain 添加域
func (cas *CasbinApi) AddDomain(c *gin.Context) {
	var addDomainForm forms.AddDomain
	if err := c.ShouldBindJSON(&addDomainForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	utils.SetDefault(&addDomainForm)
	role, err := casbinService.AddDomain(&addDomainForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(role, c)
}

// DeleteDomain 删除域
func (cas *CasbinApi) DeleteDomain(c *gin.Context) {
	var form forms.DeleteDomain
	if err := c.ShouldBindUri(&form); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := casbinService.DeleteDomain(&form)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// BatchEnforce 批量校验权限，当有一条不通过时校验就不通过
func (cas *CasbinApi) BatchEnforce(c *gin.Context) {
	var batchEnforceForm forms.BatchEnforce
	if err := c.ShouldBindJSON(&batchEnforceForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	results, err := casbinService.BatchEnforce(&batchEnforceForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(results, "ok", c)
}

// GetRolesForUserInDomain 查询用户在域上的所有角色
func (cas *CasbinApi) GetRolesForUserInDomain(c *gin.Context) {
	var userDomainForm forms.UserInDomain
	if err := c.ShouldBind(&userDomainForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	if err := utils.Verify(userDomainForm, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	results := casbinService.GetRolesForUserInDomain(&userDomainForm)
	response.OkWithData(results, c)
}

// GetPermissionsForSubInDomain 查询域角色或用户权限
func (cas *CasbinApi) GetPermissionsForSubInDomain(c *gin.Context) {
	var form forms.SubInDomain
	if err := c.ShouldBindQuery(&form); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	res := casbinService.GetPermissionsForSubInDomain(&form)
	response.OkWithData(res, c)
}

// AddPermissionsForSubInDomain 添加域角色或用户权限
func (cas *CasbinApi) AddPermissionsForSubInDomain(c *gin.Context) {
	var addPolicyForm forms.Policies
	if err := c.ShouldBindJSON(&addPolicyForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := casbinService.AddPermissionsForSubInDomain(&addPolicyForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// UpdatePermissionsForSubInDomain 更新域角色或用户权限
func (cas *CasbinApi) UpdatePermissionsForSubInDomain(c *gin.Context) {
	var form forms.UpdatePolicies
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := casbinService.UpdatePermissionsForSubInDomain(&form)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
