package controller

import (
	"gin-demo/common"
	"gin-demo/dto"
	"gin-demo/model"
	"gin-demo/response"
	"gin-demo/util"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	username := c.PostForm("username")
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	if len(mobile) != 11 {
		response.Response(c, 400, "400", nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, 400, "400", nil, "密码长度不能小于6位")
		return
	}

	if len(username) == 0 {
		username = util.RandomString(10)
	}

	log.Println(username, mobile, password)

	if isPhoneExist(DB, mobile) {
		response.Response(c, 400, "400", nil, "用户已存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, 500, "500", nil, "加密错误")
		return
	}

	newUser := model.User{
		UserName: username,
		Mobile:   mobile,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)

	response.Success(c, nil, "注册成功")
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	DB := common.GetDB()

	var params LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Response(c, 400, "400", nil, "参数错误")
		return
	}

	mobile := params.Username
	password := params.Password

	if len(mobile) != 11 {
		response.Response(c, 400, "400", nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, 400, "400", nil, "密码长度不能小于6位")
		return
	}

	var user model.User
	DB.Where("mobile = ?", mobile).First(&user)
	if user.ID == 0 {
		response.Response(c, 400, "400", nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, 400, "400", nil, "密码错误")
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(c, 500, "500", nil, "系统错误")
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": dto.TOUserDto(user.(model.User))}, "成功")
}

func isPhoneExist(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	return user.ID != 0
}
