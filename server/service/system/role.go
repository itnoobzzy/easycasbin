package system

import (
	"errors"
	"strings"

	"gorm.io/gorm/clause"

	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
)

type RoleService struct{}

// AddRole 添加域角色
func (rs *RoleService) AddRole(form *forms.DomainRole) (role models.Role, err error) {
	role.Name = form.RoleName
	role.Domain = form.DomainName
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

// GetDomainSubsForRole 获取指定域角色下所有用户
func (rs *RoleService) GetDomainSubsForRole(form *forms.DomainRole) (roles interface{}) {
	role := "role:" + form.RoleName
	domain := "domain:" + form.DomainName
	subs, _ := global.CASBIN_ENFORCER.GetImplicitUsersForRole(role, domain)
	return subs
}

// AddRoleForSubInDomain 为用户添加域角色或者为角色继承另一个角色权限
func (rs *RoleService) AddRoleForSubInDomain(form *forms.RoleForSubInDomain) error {
	var (
		role models.Role
	)

	err := global.CASBIN_DB.Where("name = ? and domain = ?",
		strings.Split(form.Role, ":")[1], strings.Split(form.Domain, ":")[1]).First(&role).Error
	if err != nil {
		return errors.New("该角色不存在！")
	}

	if _, err = global.CASBIN_ENFORCER.AddRoleForUserInDomain(form.Sub, form.Role, form.Domain); err != nil {
		return errors.New("添加域角色失败！")
	}
	return nil
}

// DeleteRoleForSubInDomain 删除角色下subject（鉴权主体）
func (rs *RoleService) DeleteRoleForSubInDomain(form *forms.RoleForSubInDomain) error {
	if _, err := global.CASBIN_ENFORCER.DeleteRoleForUserInDomain(form.Sub, form.Role, form.Domain); err != nil {
		return errors.New("删除失败！")
	}
	return nil
}
