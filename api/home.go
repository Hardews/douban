package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MovieBaseInfo(c *gin.Context) {
	err, infos := service.GetMovieBaseInfo()
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get base info failed,err :", err)
		return
	}
	for i, _ := range infos {
		c.JSON(200, gin.H{
			"name":     infos[i].Name,
			"score":    infos[i].Score,
			"year":     infos[i].Year,
			"time":     infos[i].Time,
			"area":     infos[i].Area,
			"director": infos[i].Director,
			"starring": infos[i].Starring,
		})
	}
}

func MovieInfo(c *gin.Context) {

}
