package dto

import "gin-demo/model"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func TOUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
