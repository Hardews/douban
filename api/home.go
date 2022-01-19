package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	Want := c.Param("find")

	err, infos := service.Find(Want)
	if err != nil {
		fmt.Println("find movie failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	for i, _ := range infos {
		c.JSON(200, gin.H{
			"name":      infos[i].Name,
			"score":     infos[i].Score,
			"year":      infos[i].Year,
			"area":      infos[i].Area,
			"director":  infos[i].Director,
			"starring":  infos[i].Starring,
			"Introduce": infos[i].Introduce,
			"Language":  infos[i].Language,
			"types":     infos[i].Types,
		})
	}
}

func FindWithCategory(c *gin.Context) {
	category := c.Param("category")
	err, infos := service.FindWithCategory(category)
	if err != nil {
		fmt.Println("get category movie failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	for i, _ := range infos {
		c.JSON(200, gin.H{
			"name":      infos[i].Name,
			"score":     infos[i].Score,
			"year":      infos[i].Year,
			"area":      infos[i].Area,
			"director":  infos[i].Director,
			"starring":  infos[i].Starring,
			"Introduce": infos[i].Introduce,
			"Language":  infos[i].Language,
			"types":     infos[i].Types,
		})
	}

}

func Recommend(c *gin.Context) {

}
