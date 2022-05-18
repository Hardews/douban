package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

func NewMovie(c *gin.Context) {
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

	err, movieNum := service.NewMovie(movie)
	if err != nil {
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, movieNum)
}
