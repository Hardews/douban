package api

import (
	"douban/model"
	"douban/tool"
	"fmt"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/client"
	"github.com/MashiroC/begonia/app/option"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

var (
	admFunc client.RemoteFunSync // 远程调用函数
)

func Init() {
	c := begonia.NewClient(option.Addr(":12306"))

	// 获取管理员服务
	s, err := c.Service("AdmServer")
	if err != nil {
		panic(err)
	}

	// 获取一个远程函数的同步调用
	admFunc, err = s.FuncSync("NewMovie")
	if err != nil {
		panic(err)
	}
}

func NewMovie(c *gin.Context) {
	// 初始化
	Init()

	var movie model.MovieInfo
	movie.Name, _ = c.GetPostForm("movieName")
	movie.Score, _ = c.GetPostForm("score")
	movie.Area, _ = c.GetPostForm("area")
	movie.Types, _ = c.GetPostForm("types")
	movie.Introduce, _ = c.GetPostForm("introduce")
	movie.Year, _ = c.GetPostForm("year")
	movie.Img, _ = c.GetPostForm("img")

	//借鉴B站教程
	file, err := c.FormFile("img")
	if err != nil {
		fmt.Println("get file failed,err:", err)
		tool.RespErrorWithDate(c, "上传失败")
		return
	}

	if file.Size > 1024*1024*50 {
		tool.RespErrorWithDate(c, "文件大小不合适")

		return
	}

	fileSuffix := path.Ext(file.Filename)
	if !(fileSuffix == ".jpg" || fileSuffix == ".png") {
		tool.RespErrorWithDate(c, "文件格式错误")
		return
	}

	//保存到本地
	fileName := "./movieFile/" + strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("保存错误,err", err)
		return
	}

	flag, movieNum := admFunc(movie)
	if !flag.(bool) {
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, movieNum)
}
