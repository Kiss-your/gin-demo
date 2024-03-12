package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"type:varchar(20);not null"`
	Password      string `gorm:"type:varchar(255);not null"`
	Mobile        string `gorm:"type:varchar(11);not null;unique"`
	SmsCode       string `gorm:"type:varchar(6)"`
	SmsCodeExpire string `gorm:"type:datetime"`
	AttempCount   int    `gorm:"type:int;default:0"`
	LockoutUntil  string `gorm:"type:datetime"`
}
