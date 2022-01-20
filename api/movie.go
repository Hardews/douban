package api

import (
	"database/sql"
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

func ShortComment(c *gin.Context) {
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	commentTxt := c.PostForm("ShortComment")

	movieNum, err := strconv.Atoi(num)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	flag := service.CheckSensitiveWords(commentTxt)
	if !flag {
		tool.RespErrorWithDate(c, "短评含有敏感词汇")
		return
	}
	res := service.CheckTxtLengthS(commentTxt)
	if !res {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	err = service.Comment(commentTxt, username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("comment failed,err:", err)
		return
	}

	tool.RespSuccessful(c)

}

func LongComment(c *gin.Context) {
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	commentTxt := c.PostForm("LongComment")

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

	movieNum, _ := strconv.Atoi(num)
	err := service.CommentMovie(commentTxt, username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("comment failed,err:", err)
		return
	}

	tool.RespSuccessful(c)

}

func GetLongComment(c *gin.Context) {
	num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(num)
	err, comments := service.GetComment(movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无影评")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get comment failed,err:", err)
		return
	}

	for i, _ := range comments {
		c.JSON(200, gin.H{
			"username": comments[i].Username,
			"txt":      comments[i].Txt,
			"time":     comments[i].Time,
		})
	}
}

func GetShortComment(c *gin.Context) {
	num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(num)
	err, comments := service.GetMovieComment(movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无短评")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get comment failed,err:", err)
		return
	}

	for i, _ := range comments {
		c.JSON(200, gin.H{
			"username": comments[i].Username,
			"txt":      comments[i].Txt,
			"time":     comments[i].Time,
		})
	}
}
