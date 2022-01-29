package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"movie/service"
	"movie/tool"
)

func selfPage(c *gin.Context) {
	IUsername, _ := c.Get("username")
	username := IUsername.(string)
	introduce, err := service.SelectSelfIntroduce(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "您还没有自我介绍..")
		}
		tool.RespInternalError(c)
	}
	tool.RespSuccessfulWithData(c, introduce)
	myCritics, err := service.MyCritics(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "您还没有写过影评...")
		}
		tool.RespInternalError(c)
	}
	tool.RespSuccessfulWithData(c, myCritics)
	movie, err := service.MyWantSee()
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "您还没有想看的电影...")
		}
		tool.RespInternalError(c)
	}
	tool.RespSuccessfulWithData(c, movie)
}
