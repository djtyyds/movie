package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	en := gin.Default()
	en.POST("/register", register) // 注册
	en.POST("/login", login)       // 登录
	userGroup := en.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/SignOut", SignOut)     // 退出登录
		userGroup.POST("/myCritics", MyCritics) //我的评论
		userGroup.POST("/selfPage", selfPage)   // 个人页面
	}
	movieGroup := en.Group("/movie")
	{
		movieGroup.POST("/homepage", homepage)       // 主页
		movieGroup.POST("/addPost", shortComment)    // 添加短评
		movieGroup.POST("/addCritics", filmCritics)  // 添加影评
		movieGroup.POST("/findByClass", findByClass) // 分类找电影
	}
	actorGroup := en.Group("/actor")
	{
		actorGroup.POST("/actor", actor) // 影人界面
	}
}
