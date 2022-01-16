package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	en := gin.Default()
	en.POST("/register", register) //注册
	en.POST("/login", login)       //登录
	userGroup := en.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/SignOut", SignOut) //修改密码
	}
}
