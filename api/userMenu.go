package api

import (
	"database/sql"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WantSee(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	err, wantSee := service.GetUserWantSee(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时无想看内容")
		}
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}

	err = tool.WantSee(c, wantSee, username)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}
}

func Seen(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	err, seen := service.GetUserSeen(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时无看过内容")
		}
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}

	err = tool.Seen(c, seen, username)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get wantSee failed ,err:", err)
		return
	}
}

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
		"introduce": user.Introduce,
	})
}
