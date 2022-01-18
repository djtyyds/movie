package api

import "github.com/gin-gonic/gin"

func homepage(c *gin.Context) {
	movieName := c.PostForm("movieName")
	movie :=
}
