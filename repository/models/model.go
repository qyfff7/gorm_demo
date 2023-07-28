package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"not null;comment:用户名称;size:20"`
	Password string `gorm:"not null;comment:用户密码;size:20"`
	Age      uint   `gorm:"comment:用户年龄"`
}
