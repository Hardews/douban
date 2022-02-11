package api

import (
	"douban/service"
	"douban/tool"

	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateShortComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	txt := c.PostForm("comment")

	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err, flag, _ := service.SelectComment(username, movieNum, 2, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select comment failed,err:", err)
		return
	}
	if flag {
		tool.RespErrorWithDate(c, "无评论")
		return
	}

	err = service.UpdateComment(username, txt, movieNum, 2, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update short comment failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功!")
}

func UpdateLongComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	txt := c.PostForm("comment")

	num := c.Param("movieNum")
	movieNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("translate num failed , err:", err)
		tool.RespInternetError(c)
		return
	}

	err, flag, _ := service.SelectComment(username, movieNum, 1, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("select comment failed,err:", err)
		return
	}
	if flag {
		tool.RespErrorWithDate(c, "无评论")
		return
	}
	err = service.UpdateComment(username, txt, movieNum, 1, 0)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("update Long comment failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功!")
}

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

	err, flag := service.DeleteShortComment(username, movieNum)
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
	label := c.PostForm("label") //用户存储标签
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
	label := c.PostForm("label") //用户存储标签
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
		if err == sql.ErrNoRows {
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

	err, flag, _ = service.SelectComment(username, movieNum, 2, 0)
	if err != nil {
		fmt.Println("select short comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有影评，是否更新")
		num1 := c.PostForm("choose")
		choose, _ := strconv.Atoi(num1)

		switch choose {
		case 1:
			return
		case 2:
			err = service.UpdateComment(username, commentTxt, movieNum, 2, 0)
			if err != nil {
				fmt.Println("update short comment failed,err:", err)
				tool.RespInternetError(c)
				return
			}
		}
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
	commentTopic := c.PostForm("topic")

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

	err, flag, _ := service.SelectComment(username, movieNum, 1, 0)
	if err != nil {
		fmt.Println("select long comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已有影评")
		return
	}

	err = service.CommentMovie(commentTxt, username, commentTopic, movieNum)
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

	tool.RespSuccessfulWithDate(c, comments)
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

	tool.RespSuccessfulWithDate(c, comments)
}
