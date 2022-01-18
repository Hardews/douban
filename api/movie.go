package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAMovieInfo(c *gin.Context) {
	movieName := c.PostForm("movieName")
	err, movieNum := service.FindMovieNumByName(movieName)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get movieNum failed,err:", err)
		return
	}

	err, movieInfo := service.GetAMovieInfo(movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get a movie info failed , err ", err)
		return
	}

	c.JSON(200, gin.H{
		"name":       movieInfo.Name,
		"score":      movieInfo.Score,
		"year":       movieInfo.Year,
		"time":       movieInfo.Time,
		"area":       movieInfo.Area,
		"director":   movieInfo.Director,
		"starring":   movieInfo.Starring,
		"Writer":     movieInfo.Writer,
		"CommentNum": movieInfo.CommentNum,
		"Introduce":  movieInfo.Introduce,
		"Language":   movieInfo.Language,
		"WantSee":    movieInfo.WantSee,
		"Seen":       movieInfo.Seen,
	})
}
