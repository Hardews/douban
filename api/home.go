package api

import (
	"database/sql"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func Find(c *gin.Context) {
	Want := c.Param("find")

	err, infos := service.Find(Want)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "抱歉，暂时没有您想要的电影")
			return
		}
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
			"otherName": infos[i].OtherName,
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
	rand.Seed(time.Now().UnixMicro())
	var nums []int
	for i := 0; i < 7; i++ {
		num := rand.Intn(100) + 1
		nums = append(nums, num)
	}

	for _, num := range nums {
		err, infos := service.GetAMovieInfo(num)
		if err != nil {
			tool.RespInternetError(c)
			fmt.Println("get recommend movie failed,err:", err)
			return
		}
		c.JSON(200, gin.H{
			"name":      infos.Name,
			"otherName": infos.OtherName,
			"score":     infos.Score,
			"year":      infos.Year,
			"area":      infos.Area,
			"director":  infos.Director,
			"starring":  infos.Starring,
			"Introduce": infos.Introduce,
			"Language":  infos.Language,
			"types":     infos.Types,
		})
	}

}
