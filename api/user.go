package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie/model"
	"movie/service"
	"movie/tool"
)

func register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	user := model.User{
		Name:     username,
		Password: password,
	}
	flag, err := service.IsRepeatName(user.Name)
	if err != nil {
		fmt.Println("judge repeat name err:", err)
		tool.RespInternalError(c)
		return
	}
	if flag {
		tool.RespErrorWithData(c, "该用户名已存在...")
		return
	}
	bool := service.IsPasswordPlausible(user)
	if !bool {
		tool.RespErrorWithData(c, "密码必须大于8位数")
		return
	}
	err = service.Register(user)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, err := service.IsPasswordRight(username, password)
	if err != nil {
		fmt.Println("judge password err:", err)
		tool.RespInternalError(c)
		return
	}
	if !flag {
		tool.RespErrorWithData(c, "你的密码有误...")
		return
	}
	c.SetCookie("username", username, 600, "/", "", false, false)
	tool.RespSuccessful(c)
}

func SignOut(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", "", false, false)
	tool.RespSuccessful(c)
}
