package user_ser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"xingchen/middleware"
	"xingchen/model"
)

// AccountRegister
// @Summary      账号注册
// @Description  通过账号密码注册
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param        username    query     string  true  "用户账号"
// @Param        password    query     string  true  "用户密码"
// @Router       /v1/user/register [post]
func AccountRegister(c *gin.Context) {
	u := StructUser{}
	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "massage": err.Error()})
		return
	}

	user := model.XcUser{
		Username: u.Username,
		Password: u.Password,
		Status:   1,
		VipTime:  strconv.FormatInt(time.Now().Unix(), 10),
	}

	tx := model.QueryUserExistsByUsername(u.Username)
	err = tx.First(&model.XcUser{}).Error
	if err != nil {
		if err.Error() == "record not found" {
			err := model.CreateUser(&user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "注册失败", "data": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "注册成功"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "用户存在"})
}

// AccountLogin
// @Summary      账号登录
// @Description  通过账号密码登录
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param        username    query     string  true  "用户账号"
// @Param        password    query     string  true  "用户密码"
// @Router       /v1/user/accountLogin [post]
func AccountLogin(c *gin.Context) {
	u := StructUser{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "massage": err.Error()})
		return
	}
	user := model.XcUser{}
	tx := model.QueryUserExistsByUsername(u.Username)
	err = tx.First(&user).Error

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"code": 201, "massage": "用户不存在"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 202, "massage": err.Error()})
		return
	}

	if user.Password == u.Password {
		token, e := middleware.CreateJwtToken(&user)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{"code": 201, "massage": "登录异常", "err": err})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{
				"code":         200,
				"massage":      "登录成功",
				"nickname":     user.Nickname,
				"vip_exp_time": user.VipTime,
				"is_vip":       user.IsVip,
				"head_img":     user.HeadImg,
				"sex":          user.Sex,
				"Balance":      user.Balance,
				"token":        token},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 201, "massage": "密码错误"})
}
