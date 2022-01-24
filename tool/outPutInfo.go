package tool

import (
	"douban/modle"
	"douban/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RespMovieInfo(ctx *gin.Context, movieInfo modle.MovieInfo) {
	ctx.JSON(200, gin.H{
		"name":       movieInfo.Name,
		"otherName":  movieInfo.OtherName,
		"score":      movieInfo.Score,
		"year":       movieInfo.Year,
		"time":       movieInfo.Time + "分钟",
		"area":       movieInfo.Area,
		"director":   movieInfo.Director,
		"starring":   movieInfo.Starring,
		"CommentNum": movieInfo.CommentNum,
		"Introduce":  movieInfo.Introduce,
		"Language":   movieInfo.Language,
		"WantSee":    movieInfo.WantSee,
		"Seen":       movieInfo.Seen,
		"types":      movieInfo.Types,
	})
}

func RespMovieInfos(ctx *gin.Context, infos []modle.MovieInfo) {
	for i, _ := range infos {
		ctx.JSON(200, gin.H{
			"name":       infos[i].Name,
			"otherName":  infos[i].OtherName,
			"score":      infos[i].Score,
			"year":       infos[i].Year,
			"time":       infos[i].Time + "分钟",
			"area":       infos[i].Area,
			"director":   infos[i].Director,
			"starring":   infos[i].Starring,
			"commentNum": infos[i].CommentNum,
			"Introduce":  infos[i].Introduce,
			"Language":   infos[i].Language,
			"wantSee":    infos[i].WantSee,
			"Seen":       infos[i].Seen,
			"types":      infos[i].Types,
		})
	}
}

func WantSee(c *gin.Context, infos []modle.UserHistory, username string) error {
	var num = 0
	for i, _ := range infos {
		num++
		err, movieInfo := service.GetAMovieInfo(infos[i].MovieNum)
		if err != nil {
			return err
		}
		RespMovieInfo(c, movieInfo)
		c.JSON(200, gin.H{
			"username": username,
			"comment":  infos[i].Comment,
			"movieNum": infos[i].MovieNum,
			"label":    infos[i].Label,
		})
	}
	n := strconv.Itoa(num)
	c.JSON(200, gin.H{
		"num": n + "部看过",
	})
	return nil
}

func Seen(c *gin.Context, infos []modle.UserHistory, username string) error {
	var num = 0
	for i, _ := range infos {
		num++
		err, movieInfo := service.GetAMovieInfo(infos[i].MovieNum)
		if err != nil {
			return err
		}
		RespMovieInfo(c, movieInfo)
		c.JSON(200, gin.H{
			"username": username,
			"comment":  infos[i].Comment,
			"movieNum": infos[i].MovieNum,
			"label":    infos[i].Label,
		})
	}
	n := strconv.Itoa(num)
	c.JSON(200, gin.H{
		"num": n + "部看过",
	})
	return nil
}
