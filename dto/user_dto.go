package dto

import "gin-demo/model"

type UserDto struct {
	UserName string `json:"username"`
	Mobile   string `json:"mobile"`
}

func TOUserDto(user model.User) UserDto {
	return UserDto{
		UserName: user.UserName,
		Mobile:   user.Mobile,
	}
}
