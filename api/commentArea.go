package api

import (
	"database/sql"
	"douban/modle"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func doNotLikeComment(c *gin.Context) {

}

func doNotLike(c *gin.Context) {

}

func deleteComment(c *gin.Context) {

}

func deleteCommentArea(c *gin.Context) {

}

func GetCommentArea(c *gin.Context) {
	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("shift num failed,err =", err)
		tool.RespInternetError(c)
		return
	}

	err, commentArea := service.GetCommentArea(movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无话题")
			return
		}
		fmt.Println("get comment area failed ,err:", err)
		tool.RespInternetError(c)
		return
	}

	for i, _ := range commentArea {
		c.JSON(200, gin.H{
			"username":   commentArea[i].Username,
			"topic":      commentArea[i].Topic,
			"time":       commentArea[i].Time,
			"commentNum": commentArea[i].CommentNum,
			"likeNum":    commentArea[i].LikeNum,
		})

		err, comment := service.GetCommentByNum(movieNum, commentArea[i].Num)
		if err != nil {
			fmt.Println("get comment failed ,err:", err)
			tool.RespInternetError(c)
			return
		}
		c.JSON(200, gin.H{
			"username": comment[i].Username,
			"comment":  comment[i].Comment,
			"time":     comment[i].Time,
			"likeNum":  comment[i].LikeNum,
		})
	}

}

func GiveTopicLike(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)

	Num1 := c.Param("movieNum")
	Num2 := c.Param("num")
	MovieNum, err := strconv.Atoi(Num1)
	areaNum, err := strconv.Atoi(Num2)
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}

	err, flag := service.GiveTopicLike(Username, MovieNum, areaNum)
	if err != nil {
		fmt.Println("give like failed,err,", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "您已经点过赞了！")
		return
	}
	tool.RespSuccessfulWithDate(c, "点赞成功!")
}

func GiveComment(c *gin.Context) {
	var comment modle.CommentArea
	var err error
	iUsername, _ := c.Get("username")
	comment.Username = iUsername.(string)

	Num1 := c.Param("movieNum")
	Num2 := c.Param("num")
	comment.MovieNum, err = strconv.Atoi(Num1)
	comment.Num, err = strconv.Atoi(Num2)
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}
	comment.Comment = c.PostForm("comment")

	err = service.GiveComment(comment)
	if err != nil {
		fmt.Println("give comment failed ,err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "评论成功")
}

func GiveCommentLike(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)

	Num1 := c.Param("movieNum")
	Num2 := c.Param("num")
	Num3 := c.Param("commentNum")
	MovieNum, err := strconv.Atoi(Num1)
	areaNum, err := strconv.Atoi(Num2)
	commentNum, err := strconv.Atoi(Num3)
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}

	err, flag := service.GiveCommentLike(Username, MovieNum, areaNum, commentNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("give comment like failed ,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "您已经点过赞了")
		return
	}
	tool.RespSuccessfulWithDate(c, "点赞成功！")
}

func SetCommentArea(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	Num := c.Param("movieNum")
	topic := c.PostForm("topic")

	movieNum, err := strconv.Atoi(Num)
	if err != nil {
		fmt.Println("shift num failed,err =", err)
		tool.RespInternetError(c)
		return
	}
	err = service.SetCommentArea(username, topic, movieNum)
	if err != nil {
		fmt.Println("set comment area failed,err =", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "成功设置讨论话题")
}
