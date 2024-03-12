package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleId       int `gorm:"not null"`
	PermissionId int `gorm:"not null"`
}
