package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"errors"
	"gorm.io/gorm/clause"
)

type RoleService struct{}

// AddRole 添加域角色
func (rs *RoleService) AddRole(form *forms.AddDomainRole) (role models.Role, err error) {
	role.Name = form.Name
	role.Domain = form.Domain
	if err = global.CASBIN_DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"deleted_at": nil}),
	}).Create(&role).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// UpdateRoleInfo 更新域角色信息
func (rs *RoleService) UpdateRoleInfo(form *forms.UpdateDomainRole) (role models.Role, err error) {
	// 判断对应域下是否已经存在对应角色名
	result := global.CASBIN_DB.Where(&models.Role{Name: form.NewRoleName, Domain: form.DomainName}).First(&role)
	if result.RowsAffected > 0 {
		return role, errors.New("将修改角色名称已存在！")
	}
	// 判断域下是否存在对应角色
	result = global.CASBIN_DB.Where(&models.Role{
		Name:   form.RoleName,
		Domain: form.DomainName,
	}).First(&role)
	if result.RowsAffected == 0 {
		return role, errors.New("该角色不存在！")
	}
	// 修改该角色信息，并且修改对应域下的关于该角色的权限信息
	result.Update("name", form.NewRoleName)
	// 更新权限表中域下角色名称
	rules := global.CASBIN_DB.Where(&models.CasbinRule{
		Ptype: "g",
		V1:    "role:" + form.RoleName,
		V2:    "domain:" + form.DomainName,
	}).Find(&[]models.CasbinRule{})
	rules.Update("v1", "role:"+form.NewRoleName)
	// 更新权限表中该角色的权限信息
	rules = global.CASBIN_DB.Where(&models.CasbinRule{
		Ptype: "p",
		V0:    "role:" + form.RoleName,
		V1:    "domain:" + form.DomainName,
	}).Find(&[]models.CasbinRule{})
	rules.Update("v0", "role:"+form.NewRoleName)
	return role, nil
}

// DeleteRole 删除对应域的角色
func (rs *RoleService) DeleteRole(form *forms.DeleteDomainRole) (err error) {
	roleName := form.RoleName
	domainName := form.DomainName
	global.CASBIN_DB.Where("domain=?", domainName).Delete(&models.Role{})
	_, err = global.CASBIN_ENFORCER.DeleteRole("role:" + roleName)
	if err != nil {
		return err
	}
	return
}

// GetDomainRoles 获取指定域下所有角色
func (rs *RoleService) GetDomainRoles(form *forms.GetAllRoles) (roles []models.Role) {
	domain := form.Domain
	global.CASBIN_DB.Table("role").Where("domain = ?", domain).Find(&roles)
	return roles
}
