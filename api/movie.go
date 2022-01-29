package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"movie/model"
	"movie/service"
	"movie/tool"
	"time"
)

func shortComment(c *gin.Context) {
	IUsername, _ := c.Get("username")
	username := IUsername.(string)
	txt := c.PostForm("txt")
	shortComment := model.ShortComment{
		Username1: username,
		Txt:       txt,
		Useful1:   0,
	}
	err := service.AddShortComment(shortComment)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c)
}

func filmCritics(c *gin.Context) {
	IUsername, _ := c.Get("username")
	username := IUsername.(string)
	txt := c.PostForm("txt")
	filmCritic := model.FilmCritics{
		Username:    username,
		Critics:     txt,
		CriticsTime: time.Now(),
		Useful:      0,
		Useless:     0,
		Node:        model.Node{},
	}
	err := service.AddFilmCritics(filmCritic)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c)
}

func list(c *gin.Context) { // 排行榜
	movies, err := service.List()
	if err != nil {
		tool.RespInternalError(c)
	}
	var i int
	for i = 0; i < len(movies)-1; i++ {
		tool.RespSuccessfulWithData(c, movies)
	}
}

func findByClass(c *gin.Context) {
	IClass, _ := c.Get("class")
	class := IClass.(string)
	movies, err := service.FindBYClass(class)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "暂时没有这个类型的电影...")
			return
		}
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfulWithData(c, movies)
}
