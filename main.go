package main

import (
	"log"
	"xingchen/docs"
	"xingchen/router"
)

func main() {
	docs.SwaggerInfo.Title = "星辰cms Api"
	docs.SwaggerInfo.Description = "星辰API 接口文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := router.Router()
	err := r.Run()
	if err != nil {
		log.Printf("服务启动失败")
	}
}
