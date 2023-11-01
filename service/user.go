package service

import "ManagerApi/model"

func GetUserByUid(uid uint) model.User {
	var user model.User
	GetDB().First(&user, uid)
	return user
}
