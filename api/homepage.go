package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"movie/model"
	"movie/service"
	"movie/tool"
)

func homepage(c *gin.Context) {
	movieName := c.PostForm("movieName")
	movie, err := service.SelectMovieByName(movieName)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(c, "没有会找到相关内容....")
			return
		}
		tool.RespInternalError(c)
		return
	}
	var movies model.Movie
	movies = movie
	tool.RespSuccessfulWithData(c, movies)
}

func movieDetails(c *gin.Context) {
	movieName, ok := c.GetQuery("movieName")
	if !ok {
		tool.RespInternalError(c)
	}
	movie, err := service.SelectMovieByNameDetail(movieName)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	var movies model.Movie
	movies = movie
	tool.RespSuccessfulWithData(c, movies)
}
