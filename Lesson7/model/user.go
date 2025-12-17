package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;unique;not null;comment:用户名"`
	Password string `json:"password" gorm:"size:128;not null;comment:密码哈希"` // 不返回密码
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}
