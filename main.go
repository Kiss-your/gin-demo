package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
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
			name = RandomString(10)
		}

		log.Println(name, phone, password)

		if isPhoneExist(db, phone) {
			c.JSON(400, gin.H{
				"code": "400",
				"msg":  "手机号已经注册",
			})
			return
		}

		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}

		db.Create(&newUser)

		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(i int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, i)

	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&User{})

	return db
}
