package admin_ser

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"xingchen/config"
	"xingchen/middleware"
	"xingchen/model"
)

func CreateTable(c *gin.Context) {
	admin := model.XcAdmin{}
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "massage": err.Error()})
		return
	}
	t, err := config.IsNotLock()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "massage": err.Error()})
		return
	}
	if t {
		c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "已经安装过了"})
		return
	}
	isCreateTable(&model.XcUser{}, "xc_user")
	isCreateTable(&model.XcComic{}, "xc_comic")
	isCreateTable(&model.XcShop{}, "xc_shop")
	isCreateTable(&model.XcType{}, "xc_type")
	isCreateTable(&model.XcAdmin{}, "xc_admin")

	admin.LoginTime = strconv.FormatInt(time.Now().Unix(), 10)
	admin.Password = middleware.CreateMd5(admin.Password)
	err = model.CreateAdmin(&admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 300, "massage": "admin 创建失败"})
		return
	}
	config.CretaeLock()
	c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "安装成功"})

}

func isCreateTable(dst interface{}, tableName string) {
	if !model.DB.Migrator().HasTable(dst) {
		if model.DB.AutoMigrate(dst) != nil {
			log.Printf("%s:创建失败\n", tableName)
		} else {
			log.Printf("%s:创建成功\n", tableName)
		}
	} else {
		log.Printf("%s:已存在\n", tableName)
	}
}
