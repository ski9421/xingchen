package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
	"xingchen/config"
	"xingchen/model"
)

type UserToken struct {
	UserName string `json:"username"`
	IsVip    uint   `json:"is_vip"`
	VipTime  string `json:"vip_time"`
	Balance  uint   `json:"balance"`
	jwt.RegisteredClaims
}

// CheckXcJwt 检查校验
func CheckXcJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("token")
		if t == "" {
			c.JSON(http.StatusOK, gin.H{"code": 301, "massage": "接口异常"})
		} else {
			_, err := ParamJwtToken(t)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 302, "massage": "Token err"})
				return
			}
		}
	}
}

// CreateJwtToken 生成Jwt token
func CreateJwtToken(data *model.XcUser) (string, error) {
	tk := UserToken{
		UserName: data.Username,
		IsVip:    data.IsVip,
		VipTime:  data.VipTime,
		Balance:  data.Balance,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Wq",
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	sToken, err := token.SignedString(config.JwtKey)
	if err != nil {
		return sToken, err
	}

	return sToken, nil
}

// ParamJwtToken 解析Token
func ParamJwtToken(tk string) (*UserToken, error) {
	t, err := jwt.ParseWithClaims(tk, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if c, ok := t.Claims.(*UserToken); ok && t.Valid {
		fmt.Printf("%#v\n", c)
		return c, nil

	} else {
		return c, err
	}
}
