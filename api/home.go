package api

import (
	"douban/service"
	"douban/tool"
	"strconv"

	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	Want := c.Param("find")

	err, nums := service.Find(Want)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "抱歉，暂时没有您想要的电影")
			return
		}
		fmt.Println("find movie failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	if nums == nil {
		tool.RespErrorWithDate(c, "抱歉，暂时没有您想要的电影")
		return
	}

	for i, _ := range nums {
		err, info := service.GetAMovieInfo(nums[i])
		movieNum := strconv.Itoa(nums[i])
		info.Url = "http://49.235.99.195:8080/movieInfo/" + movieNum
		if err != nil {
			fmt.Println("find movie failed,err:", err)
			tool.RespInternetError(c)
			return
		}
		tool.RespSuccessfulWithDate(c, info)
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
