package system

import (
	"akcasbin/forms"
	"akcasbin/global"
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
func (casbinService *CasbinService) AddRoleForUserInDomain(info *forms.AddRoleForUserInDomain) error {
	if _, err := global.CASBIN_ENFORCER.AddRoleForUserInDomain(info.Username, info.Role, info.Domain); err != nil {
		return errors.New("添加域角色失败！")
	}
	return nil
}
