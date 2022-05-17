package api

import (
	"context"
	"douban/model"
	"douban/proto"
	"douban/service"
	"douban/tool"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"log"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const address = "127.0.0.1:8070"

func uploadAvatar(c *gin.Context) {
	var user model.UserMenu
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)

	// 借鉴B站教程
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
	fileName := "./uploadFile/" + strconv.FormatInt(time.Now().Unix(), 10) + user.Username + fileSuffix
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("保存错误,err", err)
		return
	}

	user.Avatar = "http://49.235.99.195:8080/pictures/" + fileName[13:]

	//上传头像信息到数据库
	err = service.UploadAvatar(user)
	if err != nil {
		tool.RespErrorWithDate(c, "上传失败")
		fmt.Println("upload avatar failed ,err :", err)
		return
	}
	tool.RespSuccessful(c)
}

func SetQuestion(c *gin.Context) {
	var user model.User
	var set model.UserEncrypted
	var res bool
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)

	// 检查输入的各项数据是否为空
	user.Password, res = c.GetPostForm("password")
	if !res {
		tool.RespErrorWithDate(c, "密码为空")
		return
	}
	set.Question, res = c.GetPostForm("question")
	if !res {
		tool.RespErrorWithDate(c, "问题为空")
		return
	}
	set.Answer, res = c.GetPostForm("answer")
	if !res {
		tool.RespErrorWithDate(c, "答案为空")
		return
	}

	// 检查输入密码是否正确
	err, flag := service.CheckPassword(user.Username, user.Password)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("check password failed,err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "密码错误")
		return
	}

	// 向数据库插入密保问题和答案
	set.Username = user.Username
	err, flag = service.SetQuestion(set)
	if err != nil {
		fmt.Println("set question failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	// 返回信息
	if !flag {
		tool.RespErrorWithDate(c, "已设置过密保")
		return
	}
	tool.RespSuccessfulWithDate(c, "设置成功")
}

func Retrieve(c *gin.Context) {
	// 找回密码
	username, _ := c.GetPostForm("username")

	// 判断输入用户名是否为空值
	if username == "" {
		tool.RespErrorWithDate(c, "用户名为空")
		return
	}

	// 检查是否设置有密保
	question, err := service.SelectQuestion(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tool.RespErrorWithDate(c, "该账号无密保，可通过申诉找回") //发我邮箱，我帮你查（滑稽
			return
		}
		fmt.Println("select question failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	// 返回问题，让用户输入问题答案
	tool.RespSuccessfulWithDate(c, question)

	// 判断输入答案是否为空值
	answer, res := c.GetPostForm("answer")
	if !res {
		tool.RespErrorWithDate(c, "答案为空")
		return
	}

	// 检查答案是否正确
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

	// 答案正确用户输入新密码
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
	var res bool

	iUsername, _ := ctx.Get("username")
	user.Username = iUsername.(string)

	user.Password, res = ctx.GetPostForm("oldPassword")
	if !res {
		tool.RespErrorWithDate(ctx, "旧密码为空")
		return
	}

	newPassword, res := ctx.GetPostForm("newPassword")
	if !res {
		tool.RespErrorWithDate(ctx, "新密码为空")
		return
	}

	err, res := service.CheckPassword(user.Username, user.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
	var res bool
	Username, res := ctx.GetPostForm("username")
	if !res {
		tool.RespErrorWithDate(ctx, "输入的账号为空")
		return
	}
	Password, res := ctx.GetPostForm("password")
	if !res {
		tool.RespErrorWithDate(ctx, "输入的密码为空")
		return
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewDoubanClient(conn)

	resp, err := c.Login(context.Background(), &proto.User{UserName: Username, PassWord: Password, Nickname: ""})
	if err != nil && !resp.OK {
		tool.RespInternetError(ctx)
		fmt.Println(err)
		return
	}
	tool.RespSuccessfulWithDate(ctx, resp.Token)

}

func Register(ctx *gin.Context) {
	Username, _ := ctx.GetPostForm("signName")
	Password, _ := ctx.GetPostForm("signPassword")
	Nickname, _ := ctx.GetPostForm("nickName")

	if Username == "" {
		tool.RespErrorWithDate(ctx, "用户名为空")
		return
	}
	if Password == "" {
		tool.RespErrorWithDate(ctx, "密码为空")
		return
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDoubanClient(conn)

	resp, err := c.Register(ctx, &proto.User{UserName: Username, PassWord: Password, Nickname: Nickname})
	if err != nil && !resp.OK {
		tool.RespInternetError(ctx)
		fmt.Println(err)
		return
	}

	tool.RespSuccessfulWithDate(ctx, resp.Token)
}
