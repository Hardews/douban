package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespInternetError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}

func RespErrorWithDate(ctx *gin.Context, Date interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"date": Date,
	})
}

func RespSuccessfulWithDate(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"date": date,
	})
}
