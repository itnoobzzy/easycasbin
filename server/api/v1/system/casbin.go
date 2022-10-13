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

// AddRoleForUserInDomain 赋予用户域角色
func (cas *CasbinApi) AddRoleForUserInDomain(c *gin.Context) {
	var addRoleForm forms.AddRoleForUserInDomain
	if err := c.ShouldBindJSON(&addRoleForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := casbinService.AddRoleForUserInDomain(&addRoleForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// AddPolicy 添加权限
func (cas *CasbinApi) AddPolicy(c *gin.Context) {
	var addPolicyForm forms.AddPolicy
	if err := c.ShouldBindJSON(&addPolicyForm); err != nil {
		response.FailWithMessage(err, c)
		return
	}
	err := casbinService.AddPolicy(&addPolicyForm)
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
