package forms

// AddRoleForUserInDomain 赋予用户对应域角色请求
type AddRoleForUserInDomain struct {
	Username string `form:"username" json:"username" binding:"required"`
	Domain   string `form:"domain" json:"domain" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}
