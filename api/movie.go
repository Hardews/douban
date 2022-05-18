package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"gorm.io/gorm"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateShortComment(c *gin.Context) {
	var short model.ShortReview
	iUsername, _ := c.Get("username")
	short.Username = iUsername.(string)

	var res bool
	short.Txt, res = c.GetPostForm("comment")
	if !res {
		tool.RespErrorWithDate(c, "影评为空")
		return
	}

	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}
	short.MovieNum = movieNum

	err, flag, _ := service.SelectComment(short.Username, short.MovieNum, 2, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select comment failed,err:", err)
		return
	}
	if flag {
		tool.RespErrorWithDate(c, "无评论")
		return
	}

	err = service.UpdateComment(short.Username, short.Txt, short.MovieNum, 2, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update short comment failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功!")
}

func UpdateLongComment(c *gin.Context) {
	var long model.ShortReview
	iUsername, _ := c.Get("username")
	long.Username = iUsername.(string)

	var res bool
	long.Txt, res = c.GetPostForm("comment")
	if !res {
		tool.RespErrorWithDate(c, "影评为空")
		return
	}

	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}
	long.MovieNum = movieNum

	err, flag, _ := service.SelectComment(long.Username, long.MovieNum, 1, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select comment failed,err:", err)
		return
	}
	if flag {
		tool.RespErrorWithDate(c, "无评论")
		return
	}
	err = service.UpdateComment(long.Username, long.Txt, long.MovieNum, 1, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update Long comment failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功!")
}

func deleteLongComment(c *gin.Context) {
	var long model.MovieReview
	iUsername, _ := c.Get("username")
	long.Username = iUsername.(string)
	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}
	long.MovieNum = movieNum

	err, flag := service.DeleteLongComment(long.Username, long.MovieNum)
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
	var short model.ShortReview
	iUsername, _ := c.Get("username")
	short.Username = iUsername.(string)
	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}
	short.MovieNum = movieNum

	err, flag := service.DeleteShortComment(short.Username, short.MovieNum)
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
	var user model.UserSeen
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Label, _ = c.GetPostForm("label") //用户存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)
	user.Num = movieNum

	err := service.DeleteSeen(user)
	if err != nil {
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功")
}

func deleteUserWantSee(c *gin.Context) {
	var user model.UserWantSee
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Label, _ = c.GetPostForm("label") //用户存储标签
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)
	user.Num = movieNum

	err := service.DeleteWantSee(user)
	if err != nil {
		return
	}
	tool.RespSuccessfulWithDate(c, "删除成功")
}

func userWantSee(c *gin.Context) {
	var user model.UserWantSee
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Txt, _ = c.GetPostForm("comment") //简短评论，想看的理由（非影评短评,可以为空值
	user.Label, _ = c.GetPostForm("label") //用户输入存储标签,可以为空值

	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)
	user.Num = movieNum

	err := service.UserWantSee(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set movie wantSee failed,err:", err)
		return
	}
	tool.RespSuccessful(c)
}

func userSeen(c *gin.Context) {
	var user model.UserSeen
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Txt, _ = c.GetPostForm("comment") //简短评论，看过之后的感想（非影评短评
	user.Label, _ = c.GetPostForm("label") //用户输入存储标签,可以为空值

	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)
	user.Num = movieNum

	err := service.UserSeen(user)
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
		if err == gorm.ErrRecordNotFound {
			tool.RespErrorWithDate(c, "无该电影的信息")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get a movie info failed , err ", err)
		return
	}

	tool.RespSuccessfulWithDate(c, movieInfo)
}

func ShortComment(c *gin.Context) {
	var user model.ShortReview
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Txt, _ = c.GetPostForm("ShortComment")

	movieNum, err := strconv.Atoi(num)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println(err)
		return
	}
	user.MovieNum = movieNum

	flag := service.CheckSensitiveWords(user.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "短评含有敏感词汇")
		return
	}
	res := service.CheckTxtLengthS(user.Txt)
	if !res {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	err, flag, _ = service.SelectComment(user.Username, user.MovieNum, 2, 0)
	if err != nil {
		fmt.Println("select short comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有影评")
		return
	}

	err = service.Comment(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("comment failed,err:", err)
		return
	}

	tool.RespSuccessful(c)
}

func LongComment(c *gin.Context) {
	var user model.MovieReview
	num := c.Param("movieNum")
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Txt, _ = c.GetPostForm("LongComment")
	user.Title, _ = c.GetPostForm("topic")

	flag := service.CheckSensitiveWords(user.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "影评评含有敏感词汇")
		return
	}
	flag = service.CheckTxtLengthL(user.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	user.MovieNum, _ = strconv.Atoi(num)

	err, flag, _ := service.SelectComment(user.Username, user.MovieNum, 1, 0)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		fmt.Println("select long comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有影评")
		return
	}

	err = service.CommentMovie(user)
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
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "无影评")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get comment failed,err:", err)
		return
	}

	tool.RespSuccessfulWithDate(c, comments)
}

func GetShortComment(c *gin.Context) {
	num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(num)
	err, comments := service.GetMovieComment(movieNum)
	if err != nil {
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "无短评")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get comment failed,err:", err)
		return
	}

	tool.RespSuccessfulWithDate(c, comments)
}
