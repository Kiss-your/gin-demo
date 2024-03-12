package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId int `gorm:"not null"`
	RoleId int `gorm:"not null"`
}
