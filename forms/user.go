package forms

// Register 注册表单
type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	NickName string `form:"nickname" json:"nickName"`
}

// Login 登录请求表单
type Login struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	AuthorityId string `json:"authority_id" form:"authorityId"`
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-" form:"id"`                    // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password" form:"password"`       // 密码
	NewPassword string `json:"newPassword" form:"newPassword"` // 新密码
}

type ChangeUserInfo struct {
	ID       uint   `gorm:"primarykey"`                                // 主键ID
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Phone    string `json:"phone"  gorm:"comment:用户手机号"`               // 用户手机号
	Email    string `json:"email"  gorm:"comment:用户邮箱"`                // 用户邮箱
}
