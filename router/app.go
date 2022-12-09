package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"xingchen/middleware"
	"xingchen/service/admin_ser"
	"xingchen/service/user_ser"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{

		// 用户操作
		user := v1.Group("/user")
		{
			//用户信息相关
			// 账号注册
			user.POST("/accountRegister", user_ser.AccountRegister)
			user.POST("/accountLogin", user_ser.AccountLogin)
			// 使用中间件
			user.GET("/date", middleware.CheckXcJwt(), func(c *gin.Context) {
				//ip := c.ClientIP()
				ip := c.Request.Header.Get("X-Forward-For")
				if ip == "127.0.0.1" || ip == "" {
					ip = c.Request.Header.Get("X-real-ip")
				}
				c.JSON(http.StatusOK, gin.H{"code": 200, "ip": ip})
			})
		}
		// 管理员操作
		admin := v1.Group("/admin")
		{
			// 创建表
			admin.POST("/createTable", admin_ser.CreateTable)
		}
		// 创建表
		install := v1.Group("/install")
		install.POST("/", admin_ser.CreateTable)

	}

	return r
}
