# 星辰CMS

### 所需要安装的库
1. Gin 安装命令
 
`go get -u github.com/gin-gonic/gin`

2. Gorm

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

3. Swagger

安装地址 `go get -u github.com/swaggo/swag/cmd/swag`

找到go mod 安装的swag 目录 `github.com/swaggo/swag@v1.8.8/cmd/swag`

将main.go 使用终端打开

输入 `go install`

然后重新打开终端 使用 `swag init`生成decs

`go get -u github.com/swaggo/gin-swagger`

`go get -u github.com/swaggo/files`

`http://127.0.0.1:8080/swagger/index.html`

4. JWT 

安装 `go get -u github.com/golang-jwt/jwt/v4`
