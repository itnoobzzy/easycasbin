package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"errors"
	"fmt"
	"gorm.io/gorm/clause"
)

type CasbinService struct{}

func (casbinService *CasbinService) GetAllSubjects() interface{} {
	allSubjects := global.CASBIN_ENFORCER.GetAllSubjects()
	fmt.Print(allSubjects)
	return allSubjects
}

// GetRolesForUserInDomain 查询用户在域上的角色
func (casbinService *CasbinService) GetRolesForUserInDomain(form *forms.UserInDomain) (roles []string) {
	res := global.CASBIN_ENFORCER.GetRolesForUserInDomain(form.Username, form.Domain)
	return res
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

// AddDomain 添加域
func (casbinService *CasbinService) AddDomain(form *forms.AddDomain) (role models.Role, err error) {
	role.Domain = form.DomainName
	role.Name = form.RoleName
	if err = global.CASBIN_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&role).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// DeleteDomain 删除域
func (casbinService *CasbinService) DeleteDomain(form *forms.DeleteDomain) (err error) {
	var role models.Role
	global.CASBIN_DB.Where("domain=?", form.DomainName).Delete(&role)
	domain := "domain:" + form.DomainName
	_, err = global.CASBIN_ENFORCER.DeleteDomains(domain)
	if err != nil {
		return err
	}
	return
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
