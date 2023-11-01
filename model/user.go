package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `form:"username" binding:"required" gorm:"user_name; uniqueIndex;primaryKey;type:varchar(15)"`
	Password  string `form:"password" binding:"required" gorm:"password; size:255; notnull"`
	Telephone string `form:"telephone" gorm:"size:11"`
}

type UserDTO struct {
	UserName  string `json:"userName"`
	Telephone string `json:"telephone"`
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{
		UserName:  user.UserName,
		Telephone: user.Telephone,
	}
}

// 登录记录
type LoginRecord struct {
	gorm.Model
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	Source string `json:"source" form:"source" binding:"required"`
}
