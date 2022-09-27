package system

import (
	"akcasbin/forms"
	"akcasbin/models/response"
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

// AddPolicy 添加ACL权限
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
