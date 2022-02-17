package api

import (
	"douban/modle"
	"douban/service"
	"douban/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewMovie(c *gin.Context) {
	var movie modle.MovieInfo
	movie.Name, _ = c.GetPostForm("movieName")
	movie.OtherName, _ = c.GetPostForm("otherName")
	movie.Score, _ = c.GetPostForm("score")
	movie.Starring, _ = c.GetPostForm("Starring")
	movie.Area, _ = c.GetPostForm("Area")
	movie.Time, _ = c.GetPostForm("Time")
	movie.Director, _ = c.GetPostForm("Director")
	movie.Types, _ = c.GetPostForm("Types")
	movie.Introduce, _ = c.GetPostForm("Introduce")
	year, _ := c.GetPostForm("Year")
	movie.Year, _ = strconv.Atoi(year)

	err, movieNum := service.NewMovie(movie)
	if err != nil {
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, movieNum)
}
