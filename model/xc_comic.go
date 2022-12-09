package model

import "gorm.io/gorm"

type XcComic struct {
	gorm.Model
	Name           string `gorm:"type:varchar(32);comment:'漫画名称'" json:"name"`               // 漫画名
	Author         string `gorm:"type:varchar(32);comment:'作者名'" json:"author"`              // 作者名
	Pic            string `gorm:"type:varchar(128);comment:'竖图封面'" json:"pic"`               // 竖版封面
	Pix            string `gorm:"type:varchar(128);comment:'横图封面'" json:"pix"`               // 横板封面
	IsPay          uint   `gorm:"type:tinyint;default:0;comment:'是否付费'" json:"is_pay"`       // 是否付费 0 否 1 付费
	TypePayment    uint   `gorm:"type:tinyint;default:0;comment:'付费类型'" json:"type_payment"` // 付费类型 0 免费 1 Vip可看 2 余额
	Cid            uint   `gorm:"type:int(4);comment:'分类ID'" json:"cid"`                     // 分类ID
	Label          string `gorm:"type:varchar(32);comment:'标签'" json:"label"`                // 标签
	Sex            uint   `gorm:"type:tinyint;default:0;comment:'性别'" json:"sex"`            // 性别 0 保密 1 男 2 女
	State          uint   `gorm:"type:tinyint;default:1;comment:'状态'" json:"state"`          // 漫画状态 0 下架 1 上架
	Status         uint   `gorm:"type:tinyint;default:0;comment:'漫画状态'" json:"status"`       // 漫画状态 0 连载 1 完结
	IsRecommend    uint   `gorm:"type:tinyint;default:0;comment:'是否推荐'" json:"is_recommend"` // 是否推荐 0 不推荐 1 推荐
	Intro          string `gorm:"type:tinytext;comment:'简介'" json:"intro"`                   //简介
	LatestChapters string `gorm:"type:tinytext;comment:'最新章节名'" json:"latest_chapters"`      // 最新章节名
	LatestCid      uint   `gorm:"type:int(10);default:0;comment:'最新章节Id'" json:"latest_cid"` // 最新章节ID
	Fav            uint   `gorm:"type:int(10);default:0;comment:'收藏数'" json:"fav"`           // 收藏数
	Day            uint   `gorm:"type:int(10);default:0;comment:'日阅读数'" json:"day"`          // 日阅读数
	Week           uint   `gorm:"type:int(10);default:0;comment:'周阅读数'" json:"week"`         // 周阅读数
	Month          uint   `gorm:"type:int(10);default:0;comment:'月阅读数'" json:"month"`        // 月阅读数
	Sum            uint   `gorm:"type:int(10);default:0;comment:'总阅读数'" json:"sum"`          // 总阅读数
}

func (table *XcComic) TableName() string {
	return "xc_comic"
}

// AddComic 添加漫画
func AddComic(comic *XcComic) error {
	return DB.Create(&comic).Error
}

// DeleteComic 删除漫画
func DeleteComic(id uint) error {
	return DB.Delete(&XcComic{}).Where("id = ?", id).Error
}

// UpdateComic 修改漫画信息
func UpdateComic(comic *XcComic) {
	//DB.Update(&XcComic{}, comic).Where("")
}
