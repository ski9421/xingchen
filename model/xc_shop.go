package model

import "gorm.io/gorm"

type XcShop struct {
	gorm.Model
	Uid       int    `gorm:"type:int;comment:'创建者ID'" json:"uid"`                         // 创建者Id
	Name      string `gorm:"type:varchar(64);comment:'商品名称'" json:"name"`                 // 商品名称
	Content   string `gorm:"type:varchar(255);comment:'商品介绍'" json:"content"`             // 商品介绍
	Price     uint   `gorm:"type:int(10);comment:'商品价格'" json:"price"`                    // 商品价格
	ValidTime int    `gorm:"type:int(10);default:0;comment:'有效时间 天为单位'" json:"validTime"` // 有效时间 天为单位
}

func (table *XcShop) TableName() string {
	return "xc_shop"
}
