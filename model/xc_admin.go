package model

import "gorm.io/gorm"

type XcAdmin struct {
	gorm.Model
	Username    string `gorm:"type:varchar(64);comment:'用户名'" json:"username" binding:"required"` // 用户名
	Password    string `gorm:"type:varchar(32);comment:'密码'" json:"password" binding:"required"`  // 密码
	Nickname    string `gorm:"type:varchar(64);default:'管理员';comment:'昵称'" json:"nickname"`       // 昵称
	Status      uint   `gorm:"type:tinyint;default:1;comment:'状态'" json:"status"`                 // 用户状态 0 封禁 1 正常
	Permissions uint   `gorm:"type:tinyint;default:0;comment:'权限'" json:"permissions"`            // 权限 0 只能添加漫画 1 可以管理漫画 2 可以管理用户 3 全部权限
	LoginTime   string `gorm:"type:varchar(11);comment:'登录时间'" json:"login_time"`
}

func (table *XcAdmin) TableName() string {
	return "xc_admin"
}

// QueryAdminIsExist 检查admin是否存在
func QueryAdminIsExist(username string) *gorm.DB {
	return DB.Model(&XcAdmin{}).Where("username = ?", username)
}

// CreateAdmin 创建admin
func CreateAdmin(admin *XcAdmin) error {
	return DB.Create(&admin).Error
}
