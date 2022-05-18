package api

import (
	"database/sql"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func UpdateArea(c *gin.Context) {
	iUsername, _ := c.Get("username")
	Username := iUsername.(string)

	txt, _ := c.GetPostForm("topic")

	if txt == "" {
		tool.RespErrorWithDate(c, "话题为空")
		return
	}

	res := service.CheckSensitiveWords(txt)
	if !res {
		tool.RespErrorWithDate(c, "话题含有敏感词汇")
		return
	}
	res = service.CheckTxtLengthL(txt)
	if !res {
		tool.RespErrorWithDate(c, "话题长度不合法")
		return
	}

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
	var comment model.Comment
	iUsername, _ := c.Get("username")
	comment.Username = iUsername.(string)

	var res bool
	comment.Txt, res = c.GetPostForm("comment")

	if !res {
		tool.RespErrorWithDate(c, "评论为空")
		return
	}
	res = service.CheckSensitiveWords(comment.Txt)
	if !res {
		tool.RespErrorWithDate(c, "评论含有敏感词汇")
		return
	}
	res = service.CheckTxtLengthL(comment.Txt)
	if !res {
		tool.RespErrorWithDate(c, "评论长度不合法")
		return
	}

	num1 := c.Param("movieNum")
	num2 := c.Param("areaNum")

	movieNum, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}
	comment.CommentId = areaNum

	err = service.UpdateComment(comment.Username, comment.Txt, movieNum, 4, areaNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update comment failed,err :", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功")
}

func doNotLikeComment(c *gin.Context) {
	var user model.CommentLike
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	num2 := c.Param("areaNum")
	num3, _ := c.GetPostForm("commentNum")

	_, err := strconv.Atoi(num2)
	commentNum, err := strconv.Atoi(num3)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}
	user.CommentId = commentNum

	err = service.DoNotLikeComment(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
	var user model.TopicLike
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	num := c.Param("areaNum")

	areaNum, err := strconv.Atoi(num)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}
	user.TopicId = areaNum

	err = service.DoNotLikeTopic(user)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此点赞内容")
			return
		}
		fmt.Println("delete like failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "")
}

func deleteComment(c *gin.Context) {
	var userOp model.Comment
	iUsername, _ := c.Get("username")
	userOp.Username = iUsername.(string)
	num1 := c.Param("movieNum")
	num2 := c.Param("areaNum")

	_, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}
	userOp.CommentId = areaNum

	err, flag := service.DeleteComment(userOp)
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
	var userOp model.CommentArea
	num1 := c.Param("movieNum")
	num2 := c.Param("areaNum")

	movieNum, err := strconv.Atoi(num1)
	areaNum, err := strconv.Atoi(num2)
	userOp.MovieNum = movieNum
	userOp.ID = uint(areaNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}

	err, flag := service.DeleteCommentArea(userOp)
	if err != nil {
		fmt.Println("delete comment area failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "无此话题")
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功!")
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
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "无话题")
			return
		}
		fmt.Println("get comment area failed ,err:", err)
		tool.RespInternetError(c)
		return
	}

	for i, _ := range commentArea {
		tool.RespSuccessfulWithDate(c, commentArea[i])
		err, comment := service.GetCommentByNum(commentArea[i].ID)
		if err != nil {
			if err == gorm.ErrEmptySlice {
				tool.RespSuccessfulWithDate(c, "无评论")
				return
			}
			fmt.Println("get comment failed ,err:", err)
			tool.RespInternetError(c)
			return
		}
		tool.RespSuccessfulWithDate(c, comment)
	}

}

func GiveTopicLike(c *gin.Context) {
	var userLike model.TopicLike
	iUsername, _ := c.Get("username")
	userLike.Username = iUsername.(string)

	Num1 := c.Param("movieNum")
	Num2 := c.Param("areaNum")
	MovieNum, err := strconv.Atoi(Num1)
	areaNum, err := strconv.Atoi(Num2)
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}
	userLike.MovieNum = MovieNum
	userLike.TopicId = areaNum

	err, flag := service.GiveTopicLike(userLike)
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
	var comment model.Comment
	var err error
	iUsername, _ := c.Get("username")
	comment.Username = iUsername.(string)

	// 只需获取讨论区相关的uid就行
	Num := c.Param("areaNum")
	comment.CommentId, err = strconv.Atoi(Num)
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}

	// 获取用户评论内容
	comment.Txt, _ = c.GetPostForm("comment")

	res := service.CheckSensitiveWords(comment.Txt)
	if !res {
		tool.RespErrorWithDate(c, "评论含有敏感词汇")
		return
	}
	res = service.CheckTxtLengthL(comment.Txt)
	if !res {
		tool.RespErrorWithDate(c, "评论长度不合法")
		return
	}

	err = service.GiveComment(comment)
	if err != nil {
		fmt.Println("give comment failed ,err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "评论成功")
}

func GiveCommentLike(c *gin.Context) {
	var like model.CommentLike
	// 获取点赞人的用户名
	iUsername, _ := c.Get("username")
	like.Username = iUsername.(string)

	// 获取电影编号和讨论区的uid
	Num1 := c.Param("movieNum")
	Num2 := c.Param("areaNum")

	MovieNum, err := strconv.Atoi(Num1)
	areaNum, err := strconv.Atoi(Num2)
	like.CommentId = areaNum
	like.MovieNum = MovieNum
	if err != nil {
		fmt.Println("shift num failed,err,", err)
		tool.RespInternetError(c)
		return
	}

	err, flag := service.GiveCommentLike(like)
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
	var area model.CommentArea
	iUsername, _ := c.Get("username")
	area.Username = iUsername.(string)
	Num := c.Param("movieNum")
	topic, res := c.GetPostForm("topic")
	if !res {
		tool.RespErrorWithDate(c, "话题为空")
		return
	}

	res = service.CheckSensitiveWords(topic)
	if !res {
		tool.RespErrorWithDate(c, "话题含有敏感词汇")
		return
	}
	res = service.CheckTxtLengthL(topic)
	if !res {
		tool.RespErrorWithDate(c, "话题长度不合法")
		return
	}

	movieNum, err := strconv.Atoi(Num)
	if err != nil {
		fmt.Println("shift num failed,err =", err)
		tool.RespInternetError(c)
		return
	}
	area.MovieNum = movieNum

	err, flag, _ := service.SelectComment(area.Username, area.MovieNum, 3, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select area failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有话题")
		return
	}

	err = service.SetCommentArea(area)
	if err != nil {
		fmt.Println("set comment area failed,err =", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "成功设置讨论话题")
}
