package controller

import (
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/util"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	if len(phone) != 11 {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "密码长度不能小于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)

	if isPhoneExist(DB, phone) {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "手机号已经注册",
		})
		return
	}

	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}

	DB.Create(&newUser)

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
