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

// GetDomainRoles 获取指定域下所有角色
func (rs *RoleService) GetDomainRoles(form *forms.GetAllRoles) (roles []models.Role) {
	domain := form.Domain
	global.CASBIN_DB.Table("role").Where("domain = ?", domain).Find(&roles)
	return roles
}
