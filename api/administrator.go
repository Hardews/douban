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
	movie.Name = c.PostForm("movieName")
	movie.OtherName = c.PostForm("otherName")
	movie.Score = c.PostForm("score")
	movie.Starring = c.PostForm("Starring")
	movie.Area = c.PostForm("Area")
	movie.Time = c.PostForm("Time")
	movie.Director = c.PostForm("Director")
	movie.Types = c.PostForm("Types")
	movie.Introduce = c.PostForm("Introduce")
	year := c.PostForm("Year")
	movie.Year, _ = strconv.Atoi(year)

	err, movieNum := service.NewMovie(movie)
	if err != nil {
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, movieNum)

}
