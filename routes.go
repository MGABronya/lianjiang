// @Title  routes
// @Description  程序的路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package main

import (
	"lianjiang/controller"
	"lianjiang/middleware"

	"github.com/gin-gonic/gin"
)

// @title    CollectRoute
// @description   给gin引擎挂上路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func CollectRoute(r *gin.Engine) *gin.Engine {
	// TODO 添加中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	// TODO 用户的注册路由
	r.POST("/regist", controller.Register)

	// TODO 用户的邮箱验证
	r.GET("/verify/:id", controller.VerifyEmail)

	// TODO 用户找回密码
	r.PUT("/security", controller.Security)

	// TODO 用户更改密码
	r.PUT("/updatepass", middleware.AuthMiddleware(), controller.UpdatePass)

	// TODO 用户的登录路由
	r.POST("/login", controller.Login)

	// TODO 用户的个人页面路由
	r.GET("/personal", middleware.AuthMiddleware(), controller.PersonalPage)

	// TODO 小时文件上传
	r.POST("/hour/upload", middleware.AuthMiddleware(), controller.HourUpload)

	// TODO 小时文件列表
	r.GET("/hours")

	// TODO 小时文件下载
	//r.GET("/hour/download")

	// TODO 请求某个站点的信息
	//r.GET("/hour/:name/:")

	// TODO

	return r
}
