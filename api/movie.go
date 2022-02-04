package api

import (
	"douban/service"
	"douban/tool"

	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteLongComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err, flag := service.DeleteLongComment(username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("delete long comment failed,err: ", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "影评不存在")
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功!")
}

func deleteShortComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err, flag := service.DeleteLongComment(username, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("delete short comment failed,err: ", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "短评不存在")
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功!")
}

func deleteUserSeen(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	label := c.Param("label") //用户存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)

	err := service.DeleteSeen(movieNum, label, username)
	if err != nil {
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功")
}

func deleteUserWantSee(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	label := c.Param("label") //用户存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)

	err := service.DeleteWantSee(movieNum, label, username)
	if err != nil {
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功")
}

func userWantSee(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	comment := c.PostForm("comment") //简短评论，为甚想看
	label := c.PostForm("label")     //用户输入存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)

	err := service.UserWantSee(username, comment, label, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set movie wantSee failed,err:", err)
		return
	}
	tool.RespSuccessful(c)
}

func userSeen(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	comment := c.PostForm("comment") //简短评论，看过之后的感想（非影评短评
	label := c.PostForm("label")     //用户输入存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)

	err := service.UserSeen(username, comment, label, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set movie wantSee failed,err:", err)
		return
	}
	tool.RespSuccessful(c)
}

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

	tool.RespMovieInfo(c, movieInfo)
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
