package common

import (
	"gin-demo/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.Permission{})
	db.AutoMigrate(&model.UserRole{})
	db.AutoMigrate(&model.RolePermission{})

	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
