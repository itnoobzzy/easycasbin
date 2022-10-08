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

// policy 单个权限
type policy struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Domain   string `form:"domain" json:"domain" default:"default"`
	Resource string `form:"resource" json:"resource" binding:"required"`
	Action   string `form:"action" json:"action" binding:"required"`
	Eft      string `form:"eft" json:"eft"`
}

// AddPolicy 添加权限
type AddPolicy struct {
	Policies []policy `form:"policies" json:"policies" binding:"required"`
}

// BatchEnforce 批量校验权限
type BatchEnforce struct {
	Policies []policy `form:"policies" json:"policies" binding:"required"`
}
