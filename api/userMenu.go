package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WantSee(c *gin.Context) {
	username := c.Param("username")
	err, wantSee := service.GetUserWantSee(username)
	if err != nil {
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "暂时无内容")
		}
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}
	tool.RespErrorWithDate(c, wantSee)
}

func Seen(c *gin.Context) {
	username := c.Param("username")
	err, seen := service.GetUserSeen(username)
	if err != nil {
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "暂时无内容")
		}
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}

	tool.RespErrorWithDate(c, seen)
}

func GetUserComment(c *gin.Context) {
	username := c.Param("username")
	err, shortComments, longComments := service.GetUserComment(username)
	if err != nil {
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "该用户暂时无评论")
			return
		}
	}
	tool.RespSuccessfulWithDate(c, longComments)
	tool.RespSuccessfulWithDate(c, shortComments)
}

func SetIntroduce(c *gin.Context) {
	var user model.UserMenu
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Introduce, _ = c.GetPostForm("introduce")

	// 检查自我介绍长度和敏感词汇
	flag := service.CheckSensitiveWords(user.Introduce)
	if !flag {
		tool.RespErrorWithDate(c, "输入的自我介绍含有敏感词汇")
		return
	}
	res := service.CheckTxtLengthS(user.Introduce)
	if !res {
		tool.RespErrorWithDate(c, "自我介绍长度不合法")
		return
	}

	err := service.SetIntroduce(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set introduce failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "设置成功")
}

func GetUserInfo(c *gin.Context) {
	var user model.UserMenu
	user.Username = c.Param("username")

	err, user := service.GetUserMenu(user.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tool.RespErrorWithDate(c, "没有该用户的信息")
			return
		}
		tool.RespInternetError(c)
		fmt.Println("get menu info failed, err :", err)
		return
	}
	tool.RespSuccessfulWithDate(c, user)
}
