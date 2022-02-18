package api

import (
	"database/sql"
	"douban/middleware"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func uploadAvatar(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	//借鉴B站教程
	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println("get file failed,err:", err)
		tool.RespErrorWithDate(c, "上传失败")
		return
	}

	if file.Size > 1024*1024*5 {
		tool.RespErrorWithDate(c, "文件大小不合适")

		return
	}

	fileSuffix := path.Ext(file.Filename)
	if !(fileSuffix == ".jpg" || fileSuffix == ".png") {
		tool.RespErrorWithDate(c, "文件格式错误")
		return
	}

	//保存到本地
	fileName := "./uploadFile/" + strconv.FormatInt(time.Now().Unix(), 10) + username + fileSuffix
	fileAddress := "/opt/gocode/src/douban" + fileName[1:]
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("保存错误,err", err)
		return
	}

	fmt.Println(fileName[13:])
	loadString := "http://49.235.99.195:8080/pictures/" + fileName[13:]

	err = service.UploadAvatar(username, loadString, fileAddress)
	if err != nil {
		tool.RespErrorWithDate(c, "上传失败")
		fmt.Println("upload avatar failed ,err :", err)
		return
	}
	tool.RespSuccessful(c)
}

func SetQuestion(c *gin.Context) {
	var user model.User
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)

	user.Password, _ = c.GetPostForm("password")
	question, _ := c.GetPostForm("question")
	answer, _ := c.GetPostForm("answer")

	if user.Password == "" {
		tool.RespErrorWithDate(c, "密码为空")
		return
	}
	if question == "" {
		tool.RespErrorWithDate(c, "问题为空")
		return
	}
	if answer == "" {
		tool.RespErrorWithDate(c, "答案为空")
		return
	}
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
	username, _ := c.GetPostForm("username")

	if username == "" {
		tool.RespErrorWithDate(c, "用户名为空")
		return
	}
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
	if answer == "" {
		tool.RespErrorWithDate(c, "答案为空")
		return
	}

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
	var user model.User
	user.Username = username
	user.Password, _ = c.GetPostForm("newPassword")
	if user.Password == "" {
		tool.RespErrorWithDate(c, "新密码为空")
		return
	}

	err = service.ChangePassword(user)
	if err != nil {
		fmt.Println("change password failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "修改成功")
}

func ChangePassword(ctx *gin.Context) {
	var user model.User
	iUsername, _ := ctx.Get("username")
	user.Username = iUsername.(string)
	fmt.Println(user.Username)
	user.Password, _ = ctx.GetPostForm("oldPassword")
	newPassword, _ := ctx.GetPostForm("newPassword")

	if user.Password == "" {
		tool.RespErrorWithDate(ctx, "旧密码为空")
		return
	}
	if newPassword == "" {
		tool.RespErrorWithDate(ctx, "新密码为空")
		return
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
		user.Password = newPassword

		res = service.CheckLength(newPassword)
		if !res {
			tool.RespErrorWithDate(ctx, "新密码长度不合法")
			return
		}

		err = service.ChangePassword(user)
		if err != nil {
			tool.RespInternetError(ctx)
			fmt.Println("change password failed,err:", err)
			return
		}

		tool.RespSuccessfulWithDate(ctx, "修改成功")
	} else {
		tool.RespErrorWithDate(ctx, "旧密码错误")
		return
	}

}

func Login(ctx *gin.Context) {
	var user model.User
	var identity = "用户"
	user.Username, _ = ctx.GetPostForm("logName")
	user.Password, _ = ctx.GetPostForm("password")

	if user.Username == "" {
		tool.RespErrorWithDate(ctx, "输入的账号为空")
		return
	}
	if user.Password == "" {
		tool.RespErrorWithDate(ctx, "输入的密码为空")
		return
	}
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
	var user model.User
	user.Username, _ = ctx.GetPostForm("signName")
	user.Password, _ = ctx.GetPostForm("signPassword")
	user.NickName, _ = ctx.GetPostForm("nickName")

	if user.Username == "" {
		tool.RespErrorWithDate(ctx, "用户名为空")
		return
	}
	if user.Password == "" {
		tool.RespErrorWithDate(ctx, "密码为空")
		return
	}

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
