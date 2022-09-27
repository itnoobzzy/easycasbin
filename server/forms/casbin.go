package forms

// AddRoleForUserInDomain 赋予用户对应域角色请求
type AddRoleForUserInDomain struct {
	Username string `form:"username" json:"username" binding:"required"`
	Domain   string `form:"domain" json:"domain" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

// AddDomainRole 添加域角色
type AddDomainRole struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Domain string `form:"domain" json:"domain" binding:"required"`
}

// AddPolicy 添加ACL 权限
type AddPolicy struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Resource string `form:"resource" json:"resource" binding:"required"`
	Access   string `form:"access" json:"access" binding:"required"`
}
