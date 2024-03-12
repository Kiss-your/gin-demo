package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique"`
	Url  string `gorm:"type:varchar(100);not null"`
}
