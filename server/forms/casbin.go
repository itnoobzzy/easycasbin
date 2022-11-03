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

// Policy 单个权限
type Policy struct {
	Name     string `form:"name" json:"name" binding:"required,my_format_check"`
	Domain   string `form:"domain" json:"domain" binding:"required,my_format_check"`
	Resource string `form:"resource" json:"resource" binding:"required"`
	Action   string `form:"action" json:"action" binding:"required"`
	Eft      string `form:"eft" json:"eft"`
}

// Policies 权限列表
type Policies struct {
	Policies []Policy `form:"policies" json:"policies" binding:"required,dive"`
}

// UpdatePolicies 权限列表
type UpdatePolicies struct {
	OldPolicies []Policy `form:"old_policies" json:"old_policies" binding:"required,dive"`
	NewPolicies []Policy `form:"new_policies" json:"new_policies" binding:"required,dive"`
}

// BatchEnforce 批量校验权限
type BatchEnforce struct {
	Policies []Policy `form:"policies" json:"policies" binding:"required,dive"`
}

// UserInDomain 域上的用户
type UserInDomain struct {
	Username string `form:"username" json:"username" binding:"required,my_format_check"`
	Domain   string `form:"domain" json:"domain" binding:"required,my_format_check"`
}
