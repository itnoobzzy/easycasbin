package forms

// AddDomainRole 添加域角色
type AddDomainRole struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Domain string `form:"domain" json:"domain" binding:"required"`
}

// GetAllRoles 查询域下所有角色
type GetAllRoles struct {
	Domain string `form:"domain,default=default" json:"domain"`
}
