package api

import (
	"douban/middleware"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.Use(middleware.Cors)

	engine.POST("/register", Register)
	engine.POST("/login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(middleware.JwtToken)
		userGroup.POST("/change", ChangePassword)
		userGroup.POST("/menu/introduce/:username", SetIntroduce)
	}

	userInfo := engine.Group("/user")
	{
		userInfo.GET("/:username/menu", GetUserInfo)          //用户的信息（包括自我介绍
		userInfo.GET("/:username/Comment", GetUserComment)    //获取用户的影评和短评
		userInfo.GET("/:username/favorites/wantSee", WantSee) //收藏夹
		userInfo.GET("/:username/favorites/Seen", Seen)
	}

	home := engine.Group("/home")
	{
		home.GET("/research=:find", Find)
		home.GET("/recommend", Recommend)
		home.GET("/:category", FindWithCategory)
	}

	movie := engine.Group("/movie")
	{
		movie.Use(middleware.JwtToken)

		movie.POST("/num=:movieNum/wantSee", userWantSee) //用户想看
		movie.DELETE("/num=:movieNum/:label/wantSee", deleteUserWantSee)
		movie.POST("/num=:movieNum/seen", userSeen) //用户看过
		movie.DELETE("/num=:movieNum/:label/seen", deleteUserSeen)

		movie.POST("/num=:movieNum/longComment", LongComment) //影评
		movie.DELETE("/num=:movieNum/longComment", deleteLongComment)
		movie.POST("/num=:movieNum/longCommentUpdate", UpdateLongComment)
		movie.POST("/num=:movieNum/shortComment", ShortComment) //短评
		movie.DELETE("/num=:movieNum/shortComment", deleteShortComment)
		movie.POST("/num=:movieNum/shortCommentUpdate", UpdateShortComment)

		movie.POST("/SetNum=:movieNum/commentArea", SetCommentArea) //发表讨论区话题
		movie.DELETE("/Num=:movieNum/commentAreaNum=:num", deleteCommentArea)
		movie.POST("/Num=:movieNum/commentArea/Update", UpdateArea)

		movie.POST("/CommentNum=:movieNum/commentAreaNum=:num", GiveComment) //评论某个话题
		movie.DELETE("/movieNum=:movieNum/commentAreaNum=:areaNum/commentNum=:num", deleteComment)
		movie.POST("/Num=:movieNum/commentAreaNum=:num/update", UpdateComment)

		movie.POST("/LikeNum=:movieNum/commentAreaNum=:num", GiveTopicLike) //给某个话题点赞
		movie.POST("/Num=:movieNum/commentAreaNum=:num", doNotLike)
		movie.POST("/LikeComment-Num=:movieNum/commentAreaNum=:num/:commentNum", GiveCommentLike) //给某个评论点赞
		movie.POST("/LikeComment-Num=:movieNum/commentAreaNum=:num", doNotLikeComment)
	}

	movieGet := engine.Group("/movie")
	{
		movieGet.GET("/num=:movieNum", GetAMovieInfo) //获取电影的信息
		movieGet.GET("/GetNum=:movieNum/shortComment", GetShortComment)
		movieGet.GET("/GetNum=:movieNum/longComment", GetLongComment)
		movieGet.GET("/commentArea/GetNum=:num", GetCommentArea) //获取讨论区信息
	}

	administrator := engine.Group("/administrator")
	{
		administrator.Use(middleware.AdministratorToken)

		administrator.POST("/setNewMovie", NewMovie)
	}

	engine.Run()
}
