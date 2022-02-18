package api

import (
	"database/sql"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WantSee(c *gin.Context) {
	username := c.Param("username")
	err, wantSee := service.GetUserWantSee(username)
	if err != nil {
		if err == sql.ErrNoRows {
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
		if err == sql.ErrNoRows {
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
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "该用户暂时无评论")
			return
		}
	}
	tool.RespSuccessfulWithDate(c, longComments)
	tool.RespSuccessfulWithDate(c, shortComments)
}

func SetIntroduce(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	introduce, _ := c.GetPostForm("introduce")

	flag := service.CheckSensitiveWords(introduce)
	if !flag {
		tool.RespErrorWithDate(c, "输入的自我介绍含有敏感词汇")
		return
	}
	res := service.CheckTxtLengthS(introduce)
	if !res {
		tool.RespErrorWithDate(c, "自我介绍长度不合法")
		return
	}

	err := service.SetIntroduce(username, introduce)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("set introduce failed,err:", err)
		return
	}
	tool.RespSuccessfulWithDate(c, "设置成功")
}

func GetUserInfo(c *gin.Context) {
	username := c.Param("username")

	err, user := service.GetUserMenu(username)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get menu info failed, err :", err)
		return
	}
	c.JSON(200, gin.H{
		"username":  username,
		"nickName":  user.NickName,
		"introduce": user.Introduce,
		"img":       user.Img,
		"address":   user.ImgAddress,
	})
}
