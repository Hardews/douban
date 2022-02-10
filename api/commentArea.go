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

func UpdateArea(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)

	txt := c.PostForm("comment")

	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.UpdateComment(Username, txt, movieNum, 3, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update comment failed,err :", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功")
}

func UpdateComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)

	txt := c.PostForm("comment")

	num1 := c.Param("movieNum")
	num2 := c.Param("num")

	movieNum, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.UpdateComment(Username, txt, movieNum, 4, areaNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update comment failed,err :", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功")
}

func doNotLikeComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)
	num2 := c.Param("areaNum")
	num3 := c.Param("num")

	areaNum, err := strconv.Atoi(num2)
	commentNum, err := strconv.Atoi(num3)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	err = service.DoNotLikeComment(Username, areaNum, commentNum)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此点赞内容")
			return
		}
		fmt.Println("delete like failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "取消点赞成功")
}

func doNotLike(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)
	num := c.Param("areaNum")

	areaNum, err := strconv.Atoi(num)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	err = service.DoNotLikeTopic(Username, areaNum)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此点赞内容")
			return
		}
		fmt.Println("delete like failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "取消点赞成功")
}

func deleteComment(c *gin.Context) {
	num1 := c.Param("movieNum")
	num2 := c.Param("areaNum")
	num3 := c.Param("num")

	movieNum, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	commentNum, err := strconv.Atoi(num3)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	err, flag := service.DeleteComment(movieNum, areaNum, commentNum)
	if err != nil {
		fmt.Println("delete comment failed,err :", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "没有这个评论")
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功")
}

func deleteCommentArea(c *gin.Context) {
	num1 := c.Param("movieNum")
	num2 := c.Param("areaNum")

	movieNum, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	err, flag := service.DeleteCommentArea(movieNum, areaNum)
	if err != nil {
		fmt.Println("delete comment area failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "无此讨论标题")
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功!")
}

func GetCommentArea(c *gin.Context) {
	num := c.Param("num")
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

	if commentArea == nil {
		tool.RespErrorWithDate(c, "无讨论区")
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

		if commentArea[i].CommentNum == 0 {
			tool.RespSuccessfulWithDate(c, "无评论")
			return
		}
		err, comment := service.GetCommentByNum(movieNum, commentArea[i].Num)
		if err != nil {
			fmt.Println("get comment failed ,err:", err)
			tool.RespInternetError(c)
			return
		}
		for r, _ := range comment {
			c.JSON(200, gin.H{
				"username": comment[r].Username,
				"comment":  comment[r].Comment,
				"time":     comment[r].Time,
				"likeNum":  comment[r].LikeNum,
			})
		}

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

	err, flag, _ := service.SelectComment(username, movieNum, 3, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select area failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有话题，是否重新设置（原话题所属评论会删除）")
		num2 := c.PostForm("choose")
		choose, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Println("shift num failed,err =", err)
			tool.RespInternetError(c)
			return
		}
		switch choose {
		case 1:
			return
		case 2:
			err = service.UpdateComment(username, topic, movieNum, 3, 0)
			if err != nil {
				tool.RespInternetError(c)
				fmt.Println("update comment failed,err :", err)
				return
			}
			tool.RespSuccessfulWithDate(c, "修改成功")
		}
	}

	err = service.SetCommentArea(username, topic, movieNum)
	if err != nil {
		fmt.Println("set comment area failed,err =", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "成功设置讨论话题")
}