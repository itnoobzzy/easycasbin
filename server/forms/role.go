package forms

// AddDomainRole 添加域角色
type AddDomainRole struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Domain string `form:"domain" json:"domain" binding:"required"`
}

// DeleteDomainRole 删除对应域的角色
type DeleteDomainRole struct {
	RoleName   string `binding:"required" uri:"role_name"`
	DomainName string `binding:"required" uri:"domain_name"`
}

// UpdateDomainRole 更新域角色信息
type UpdateDomainRole struct {
	RoleName    string `form:"role_name" json:"role_name" binding:"required"`
	DomainName  string `form:"domain_name" json:"domain_name" binding:"required"`
	NewRoleName string `form:"new_role_name" json:"new_role_name" binding:"required"`
}

// GetAllRoles 查询域下所有角色
type GetAllRoles struct {
	Domain string `form:"domain,default=default" json:"domain"`
}
