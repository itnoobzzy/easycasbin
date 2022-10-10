package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"gorm.io/gorm/clause"
)

type RoleService struct{}

// AddRole 添加域角色
func (rs *RoleService) AddRole(form *forms.AddDomainRole) (role models.Role, err error) {
	role.Name = form.Name
	role.Domain = form.Domain
	if err = global.CASBIN_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&role).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (rs *RoleService) GetDomainRoles(form *forms.GetAllRoles) (roles []string) {
	domain := form.Domain
	global.CASBIN_DB.Table("role").Select("name").Where("domain = ?", domain).Find(&roles)
	return roles
}
