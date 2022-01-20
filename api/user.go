package api

import (
	"douban/modle"
	"douban/service"
	"douban/tool"
	"strconv"

	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	err, shortComments, longComments := service.GetUserComment(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无评论")
			return
		}
	}

	for i, _ := range shortComments {
		c.JSON(200, gin.H{
			"movieNum": shortComments[i].MovieNum,
			"username": username,
			"txt":      shortComments[i].Txt,
			"time":     shortComments[i].Time,
		})
	}
	for i, _ := range longComments {
		c.JSON(200, gin.H{
			"movieNum": longComments[i].MovieNum,
			"username": username,
			"txt":      longComments[i].Txt,
			"time":     longComments[i].Time,
		})
	}
}

func WantSee(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	wantSee := c.PostForm("wantSee")
	Num := c.Param("movieNum")

	movieNum, _ := strconv.Atoi(Num)

	err := service.UserWantSee(username, wantSee, movieNum)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set movie wantSee failed,err:", err)
		return
	}
	tool.RespSuccessful(c)
}

func SetIntroduce(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	introduce := c.PostForm("introduce")

	AUsername := c.Param("username")
	if username != AUsername {
		tool.RespInternetError(c)
		fmt.Println("menu failed ")
		return
	}

	flag := service.CheckSensitiveWords(introduce)
	if !flag {
		tool.RespErrorWithDate(c, "自我介绍含有敏感词汇")
		return
	}
	res := service.CheckTxtLengthS(introduce)
	if !res {
		tool.RespErrorWithDate(c, "长度不合法")
		return
	}

	err := service.SetIntroduce(username, introduce)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set introduce failed,err:", err)
		return
	}
	tool.RespSuccessful(c)
}

func GetUserInfo(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	AUsername := c.Param("username")
	if username != AUsername {
		tool.RespInternetError(c)
		fmt.Println("menu failed ")
		return
	}

	err, user := service.GetUserMenu(username)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get menu info failed, err :", err)
		return
	}
	c.JSON(200, gin.H{
		"introduce":   user.Introduce,
		"wantSee":     user.WantSee,
		"filmCritics": user.FilmCritics,
		"seen":        user.Seen,
	})
}

func ChangePassword(ctx *gin.Context) {
	var user modle.User
	iUsername, _ := ctx.Get("username")
	user.Username = iUsername.(string)
	user.Password = ctx.PostForm("oldPassword")
	newPassword := ctx.PostForm("newPassword")

	err, res := service.CheckPassword(user)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(ctx, "无此账号")
			return
		}
		fmt.Println(err)
		tool.RespInternetError(ctx)
		return
	}
	if res {
		user.Password = newPassword

		res = service.CheckLength(user.Password)
		if !res {
			tool.RespErrorWithDate(ctx, "密码长度不合法")
			return
		}

		err = service.ChangePassword(user)
		if err != nil {
			tool.RespInternetError(ctx)
			fmt.Println("change password failed,err:", err)
			return
		}

		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "旧密码错误")
		return
	}

}

func Login(ctx *gin.Context) {
	var user modle.User
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	err, res := service.CheckPassword(user)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(ctx, "无此账号")
			return
		}
		fmt.Println(err)
		tool.RespInternetError(ctx)
		return
	}
	if res {
		ctx.SetCookie("user_login", user.Username, 600, "/", "", false, true)
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	}

}

func Register(ctx *gin.Context) {
	var user modle.User
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")

	res := service.CheckSensitiveWords(user.Username)
	if !res {
		tool.RespErrorWithDate(ctx, "用户名含有敏感词汇")
		return
	}

	err, flag := service.CheckUsername(user)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("check username failed, error: ", err)
		return
	}
	if flag == false {
		tool.RespErrorWithDate(ctx, "用户名已存在!")
		return
	}

	res = service.CheckLength(user.Password)
	if !res {
		tool.RespErrorWithDate(ctx, "密码长度不合法")
		return
	}

	err, user.Password = service.Encryption(user.Password)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("encryption failed , err :", err)
		return
	}

	err = service.WriteIn(user)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("register failed,err:", err)
		return
	}

	tool.RespSuccessful(ctx)

}
