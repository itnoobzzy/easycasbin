package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"gorm.io/gorm/clause"
	"sort"
	"strings"
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

// GetPermissionsForSubInDomain 查询域角色或用户权限
func (casbinService *CasbinService) GetPermissionsForSubInDomain(form *forms.SubInDomain) (results []interface{}) {
	res, _ := global.CASBIN_ENFORCER.GetImplicitPermissionsForUser(form.Sub, form.Domain)
	type ps struct {
		Resource string
		Action   string
	}
	newSet := mapset.NewSet[ps]()
	for _, p := range res {
		var policy ps
		policy.Resource = p[2]
		policy.Action = p[3]
		newSet.Add(policy)
	}
	all := newSet.ToSlice()
	for _, v := range all {
		p := ps{
			Resource: v.Resource,
			Action:   v.Action,
		}
		results = append(results, p)
	}

	return results
}

// AddPermissionsForSubInDomain 添加域角色或用户权限
func (casbinService *CasbinService) AddPermissionsForSubInDomain(form *forms.Policies) error {
	if rules, err := casbinService.checkDomains(form.Policies); err != nil {
		return err
	} else {
		// 先删除这些权限，否则如果添加的权限中有任何已经存在的，将不会执行添加动作
		if ok, _ := global.CASBIN_ENFORCER.RemovePolicies(rules); ok {
			_, err = global.CASBIN_ENFORCER.AddPolicies(rules)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// checkDomains 校验规则中的域是否存在
// 不存在返回错误，正常返回规则列表
func (casbinService *CasbinService) checkDomains(policies []forms.Policy) (rules [][]string, err error) {
	allDomains, _ := casbinService.GetAllDomains()
	sort.Strings(allDomains)
	for _, v := range policies {
		rule := make([]string, 0)
		// 默认添加至default 域，并且添加的权限是生效的
		if v.Domain == "" {
			v.Domain = "domain:default"
		}
		if v.Eft == "" {
			v.Eft = "allow"
		}
		rule = append(rule, v.Name, v.Domain, v.Resource, v.Action, v.Eft)
		rules = append(rules, rule)
		domain := strings.Split(v.Domain, ":")[1]
		idx := sort.SearchStrings(allDomains, domain)
		if in := idx < len(allDomains) && allDomains[idx] == domain; !in {
			return rules, errors.New(fmt.Sprintf("domain: %s is not exists!", domain))
		}
	}
	return rules, nil
}

// UpdatePermissionsForSubInDomain 更新域角色或用户权限
func (casbinService *CasbinService) UpdatePermissionsForSubInDomain(form *forms.UpdatePolicies) error {
	oldRules, err := casbinService.checkDomains(form.OldPolicies)
	NewRules, err := casbinService.checkDomains(form.NewPolicies)
	// 新旧规则数量必须一样
	if len(oldRules) != len(NewRules) {
		return errors.New("修改的新旧规则数量必须一致！")
	}
	if err != nil {
		return err
	}
	updated, err := global.CASBIN_ENFORCER.UpdatePolicies(oldRules, NewRules)
	if err != nil {
		return err
	}
	fmt.Print(updated)
	if !updated {
		return errors.New("修改失败，未查询到对应规则！")
	}
	return nil
}
