package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"movie/service"
	"movie/tool"
)

func actor(c *gin.Context) {
	IName, _ := c.Get("name")
	name := IName.(string)
	actor, err := service.SelectActorByName(name)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "未找到此人的信息...")
			return
		}
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfulWithData(c, actor)
}
