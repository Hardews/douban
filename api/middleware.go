package api

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context) {
	username, err := c.Cookie("user_login")
	if err != nil {
		tool.RespErrorWithDate(c, "请登陆后再进行操作")
		c.Abort()
	}
	c.Set("username", username)
	c.Next()
}
