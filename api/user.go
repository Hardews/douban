package api

import (
	"JD/modle"
	"JD/service"
	"JD/tool"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
