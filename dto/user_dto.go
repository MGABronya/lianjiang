// @Title  user_dto
// @Description  用于封装用户信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package dto

import (
	"lianjiang/model"
)

// UserDto			定义了用户的基本信息
type UserDto struct {
	Name  string `gorm:"type:varchar(20);not null"`        // 用户名称
	Email string `gorm:"type:varchar(50);not null;unique"` // 邮箱
}

// @title    ToUserDto
// @description   用户信息封装
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    user model.User       接收一个用户类
// @return   UserDto			   返回一个用户的基本信息类
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Email: user.Email,
	}
}
