// @Title  PersonalController
// @Description  该文件用于提供操作个人界面的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"lianjiang/dto"
	"lianjiang/model"
	"lianjiang/response"

	"github.com/gin-gonic/gin"
)

// @title    PersonalPage
// @description   提供用户的个人信息
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func PersonalPage(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")

	user := tuser.(model.User)

	response.Success(ctx, gin.H{"user": dto.ToUserDto(user)}, "请求成功")

}

