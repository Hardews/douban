package api

import (
	"JD/modle"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

	err = service.WriteIn(user)
	if err != nil {
		tool.RespInternetError(ctx)
		fmt.Println("register failed,err:", err)
		return
	}

	tool.RespSuccessful(ctx)

}
