package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"errors"
	"fmt"
)

type CasbinService struct{}

func (casbinService *CasbinService) GetAllSubjects() interface{} {
	allSubjects := global.CASBIN_ENFORCER.GetAllSubjects()
	fmt.Print(allSubjects)
	return allSubjects
}

// AddRoleForUserInDomain 为用户添加域角色
func (casbinService *CasbinService) AddRoleForUserInDomain(form *forms.AddRoleForUserInDomain) error {
	var role models.Role
	err := global.CASBIN_DB.Where("name = ? and domain = ?", form.Role, form.Domain).First(&role).Error
	if err != nil {
		return errors.New("请确保该域有该角色！")
	}
	if _, err = global.CASBIN_ENFORCER.AddRoleForUserInDomain(form.Username, form.Role, form.Domain); err != nil {
		return errors.New("添加域角色失败！")
	}
	return nil
}

// GetAllDomains 获取所有域
func (casbinService *CasbinService) GetAllDomains() (domains []string, err error) {
	var roleList []models.Role
	if err = global.CASBIN_DB.Distinct("Domain").Where("deleted_at is null").Find(&roleList).Error; err != nil {
		return make([]string, 0), errors.New("获取所有域信息失败!")
	}
	for _, role := range roleList {
		domains = append(domains, role.Domain)
	}
	return domains, nil
}

// AddPolicy 添加ACL权限
func (casbinService *CasbinService) AddPolicy(form *forms.AddPolicy) error {
	if _, err := global.CASBIN_ENFORCER.AddPolicy(form.Name, form.Resource, form.Access); err != nil {
		return errors.New("添加ACL权限失败!")
	}
	return nil
}
