package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	Want := c.Param("find")

	err, infos := service.Find(Want)
	if err != nil {
		if err == gorm.ErrEmptySlice {
			tool.RespErrorWithDate(c, "抱歉，暂时没有您想要的电影")
			return
		}
		fmt.Println("find movie failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, infos)
}

func FindWithCategory(c *gin.Context) {
	category := c.Param("category")
	err, infos := service.FindWithCategory(category)
	if err != nil {
		fmt.Println("get category movie failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, infos)

}

func Recommend(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
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
		tool.RespSuccessfulWithDate(c, infos)
	}
}
