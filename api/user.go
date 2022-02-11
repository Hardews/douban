package api

import (
	"database/sql"
	"douban/middleware"
	"douban/modle"
	"douban/service"
	"douban/tool"
	"strconv"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

func uploadAvatar(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println("get file failed,err:", err)
		tool.RespErrorWithDate(c, "头像上传失败")
		return
	}

	//保存到本地
	fileName := "./uploadFile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("保存错误,err", err)
		return
	}

	loadString := "http://101.201.234.29:8080/" + fileName[1:]

	err = service.UploadAvatar(username, loadString)
	if err != nil {
		tool.RespErrorWithDate(c, "上传失败")
		fmt.Println("upload avatar failed ,err :", err)
		return
	}
	tool.RespSuccessful(c)
}

func SetQuestion(c *gin.Context) {
	var user modle.User
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)

	user.Password = c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	err, flag := service.CheckPassword(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("check password failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "密码错误")
		return
	}

	err, flag = service.SetQuestion(user.Username, question, answer)
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
			tool.RespErrorWithDate(c, "该账号无密保，可通过申诉找回") //发我邮箱，我帮你查（滑稽
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
	user.NickName = ctx.PostForm("nickName")
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	fmt.Println(user.NickName)

	res := service.CheckSensitiveWords(user.Username)
	if !res {
		tool.RespErrorWithDate(ctx, "用户名含有敏感词汇")
		return
	}
	res = service.CheckSensitiveWords(user.NickName)
	if !res {
		tool.RespErrorWithDate(ctx, "昵称含有敏感词汇")
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

	//res = service.CheckLength(user.Password)
	//if !res {
	//	tool.RespErrorWithDate(ctx, "密码长度不合法")
	//	return
	//}

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
