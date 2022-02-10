package api

import (
	"database/sql"
	"douban/middleware"
	"douban/modle"
	"douban/service"
	"douban/tool"

	"fmt"

	"github.com/gin-gonic/gin"
)

func SetQuestion(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	question := c.PostForm("question")
	answer := c.PostForm("answer")

	err, flag := service.SetQuestion(username, question, answer)
	if err != nil {
		fmt.Println("set question failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "已设置过密保")
		return
	}
	tool.RespSuccessfulWithDate(c, "设置成功")
}

func Retrieve(c *gin.Context) {
	username := c.PostForm("username")

	question, err := service.SelectQuestion(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "该账号无密保，可通过申诉找回")
			return
		}
		fmt.Println("select question failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, question)
	answer := c.PostForm("answer")

	err, flag := service.CheckAnswer(username, answer)
	if err != nil {
		fmt.Println("check answer failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "答案错误！")
		return
	}
	var user modle.User
	user.Username = username
	user.Password = c.PostForm("newPassword")

	err = service.ChangePassword(user)
	if err != nil {
		fmt.Println("change password failed,err:", err)
		tool.RespInternetError(c)
		return
	}
}

func ChangePassword(ctx *gin.Context) {
	var user modle.User
	iUsername, _ := ctx.Get("username")
	user.Username = iUsername.(string)
	fmt.Println(user.Username)
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

		res = service.CheckLength(newPassword)
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
	var identity = "用户"
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")

	flag := service.CheckAdministratorUsername(user.Username)
	if flag {
		identity = "管理员"
	}

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
		token, flag := middleware.SetToken(user.Username, identity)
		if !flag {
			tool.RespInternetError(ctx)
			return
		}
		ctx.JSON(200, gin.H{
			"msg": token,
		})
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
