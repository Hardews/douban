package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAMovieInfo(c *gin.Context) {
	Num := c.Param("movieNum")

	movieNum, err := strconv.Atoi(Num)
	if err != nil {
		fmt.Println(err)
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

func Comment(c *gin.Context) {
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	commentTxt := c.PostForm("comment")

	movieNum, _ := strconv.Atoi(num)
	err := service.Comment(commentTxt, username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("comment failed,err:", err)
		return
	}

	flag := service.CheckSensitiveWords(commentTxt)
	if !flag {
		tool.RespErrorWithDate(c, "短评含有敏感词汇")
		return
	}
	flag = service.CheckTxtLengthS(commentTxt)
	if !flag {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	tool.RespSuccessful(c)

}

func CommentMovie(c *gin.Context) {
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	commentTxt := c.PostForm("comment")

	movieNum, _ := strconv.Atoi(num)
	err := service.CommentMovie(commentTxt, username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("comment failed,err:", err)
		return
	}

	flag := service.CheckSensitiveWords(commentTxt)
	if !flag {
		tool.RespErrorWithDate(c, "影评评含有敏感词汇")
		return
	}
	flag = service.CheckTxtLengthL(commentTxt)
	if !flag {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	tool.RespSuccessful(c)

}

func GetMovieComment(c *gin.Context) {

}
