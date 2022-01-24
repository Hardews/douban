package tool

import (
	"douban/modle"
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
