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

// AddPolicy 添加权限
func (casbinService *CasbinService) AddPolicy(form *forms.AddPolicy) error {
	var rules [][]string
	for _, v := range form.Policies {
		rule := make([]string, 0)
		// 默认添加至default 域，并且添加的权限是生效的
		if v.Domain == "" {
			v.Domain = "default"
		}
		if v.Eft == "" {
			v.Eft = "allow"
		}
		rule = append(rule, v.Name, v.Domain, v.Resource, v.Action, v.Eft)
		rules = append(rules, rule)
	}
	_, err := global.CASBIN_ENFORCER.AddPolicies(rules)
	if err != nil {
		return err
	}

	return nil
}

// BatchEnforce 批量校验权限
func (casbinService *CasbinService) BatchEnforce(form *forms.BatchEnforce) (results []bool, err error) {
	var rules [][]interface{}
	for _, v := range form.Policies {
		rule := make([]interface{}, 0)
		rule = append(rule, v.Name, v.Domain, v.Resource, v.Action)
		rules = append(rules, rule)
	}
	results, err = global.CASBIN_ENFORCER.BatchEnforce(rules)
	if err != nil {
		return nil, err
	}
	return results, nil
}
