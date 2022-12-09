package model

import "gorm.io/gorm"

type XcType struct {
	gorm.Model
	Name          string `gorm:"type:varchar(64);comment:'分类名称'" json:"name"`
	IsParentClass int    `gorm:"type:int(10);comment:'是否父级分类 0 1'" json:"is_parent_class"`
	Pid           int    `gorm:"type:int(10);comment:'父级分类ID'" json:"pid"`
}

func (table *XcType) TableName() string {
	return "xc_type"
}
