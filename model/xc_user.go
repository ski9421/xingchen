package model

import (
	"gorm.io/gorm"
)

type XcUser struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);comment:'用户名'" json:"username"`             // 用户名
	Password string `gorm:"type:varchar(32);comment:'密码'" json:"password"`              // 密码
	Tel      string `gorm:"type:varchar(11);comment:'手机号'" json:"tel"`                  // 手机号
	HeadImg  string `gorm:"type:varchar(255);comment:'用户头像'" json:"headImg"`            // 用户头像
	Nickname string `gorm:"type:varchar(64);default:'匿名';comment:'昵称'" json:"nickname"` // 昵称
	Status   uint   `gorm:"type:tinyint;default:1;comment:'用户状态'" json:"status"`        // 用户状态 0 封禁 1 正常
	Sex      uint   `gorm:"type:tinyint;default:0;comment:'用户性别'" json:"sex"`           // 性别 0 保密
	IsVip    uint   `gorm:"type:int(10);default:0;comment:'是否会员 0 不是'" json:"is_vip"`   //是否会员
	VipTime  string `gorm:"type:varchar(11);comment:'会员到期时间'" json:"vip_time"`          // 会员到期时间
	Balance  uint   `gorm:"type:int(10);default:0;comment:'余额'" json:"balance"`         // 余额
}

func (table *XcUser) TableName() string {
	return "xc_user"
}

// QueryUserExistsByUsername 通过用户账号查找是否存在
func QueryUserExistsByUsername(username string) *gorm.DB {
	return DB.Model(new(XcUser)).Where("username = ?", username)
}

// CreateUser 创建用户
func CreateUser(model *XcUser) error {
	return DB.Create(&model).Error
}
