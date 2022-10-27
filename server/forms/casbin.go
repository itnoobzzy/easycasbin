package forms

// AddDomain 添加域
type AddDomain struct {
	DomainName string `form:"domain_name" json:"domain_name" binding:"required"`
	RoleName   string `form:"role_name" json:"role_name" default:"default"`
}

// DeleteDomain 删除域
type DeleteDomain struct {
	DomainName string `form:"domain_name" json:"domain_name" binding:"required" uri:"domain_name"`
}

// RoleForSubInDomain 域上对应角色的subject
type RoleForSubInDomain struct {
	Sub    string `form:"sub" json:"sub" binding:"required,my_format_check"`
	Domain string `form:"domain" json:"domain" binding:"required,my_format_check"`
	Role   string `form:"role" json:"role" binding:"required,my_format_check"`
}

// SubInDomain 域上的subject
type SubInDomain struct {
	Sub    string `form:"sub" json:"sub" binding:"required,my_format_check"`
	Domain string `form:"domain" json:"domain" binding:"required,my_format_check"`
}

// policy 单个权限
type policy struct {
	Name     string `form:"name" json:"name" binding:"required,my_format_check"`
	Domain   string `form:"domain" json:"domain" binding:"required,my_format_check"`
	Resource string `form:"resource" json:"resource" binding:"required"`
	Action   string `form:"action" json:"action" binding:"required"`
	Eft      string `form:"eft" json:"eft"`
}

// AddPolicy 添加权限
type AddPolicy struct {
	Policies []*policy `form:"policies" json:"policies" binding:"required,dive"`
}

// BatchEnforce 批量校验权限
type BatchEnforce struct {
	Policies []*policy `form:"policies" json:"policies" binding:"required,dive"`
}

// UserInDomain 域上的用户
type UserInDomain struct {
	Username string `form:"username" json:"username" binding:"required,my_format_check"`
	Domain   string `form:"domain" json:"domain" binding:"required,my_format_check"`
}
