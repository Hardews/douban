package api

import (
	"douban/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.Use(middleware.Cors)

	//部署前端静态网页
	engine.Use(static.Serve("/", static.LocalFile("../static", false)))

	engine.POST("/register", Register)
	engine.POST("/login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(middleware.JwtToken)
		userGroup.POST("/uploadAvatar", uploadAvatar)
		userGroup.POST("/change", ChangePassword)
		userGroup.POST("/introduce", SetIntroduce)
		userGroup.POST("/setQuestion", SetQuestion)
	}

	userInfo := engine.Group("/user")
	{
		userInfo.POST("/retrieve", Retrieve)
		userInfo.GET("/:username/menu", GetUserInfo)       //用户的信息（包括自我介绍,头像
		userInfo.GET("/:username/Comment", GetUserComment) //获取用户的影评和短评
		userInfo.GET("/:username/wantSee", WantSee)        //收藏夹
		userInfo.GET("/:username/Seen", Seen)
	}

	home := engine.Group("/home")
	{
		home.GET("/research/:find", Find)
		home.GET("/recommend", Recommend)
		home.GET("/:category", FindWithCategory)
	}

	movie := engine.Group("/movie")
	{
		movie.Use(middleware.JwtToken)

		movieNum := movie.Group("/:movieNum")
		{
			movieNum.POST("/wantSee", userWantSee) //用户想看
			movieNum.DELETE("/wantSee", deleteUserWantSee)

			movieNum.POST("/seen", userSeen) //用户看过
			movieNum.DELETE("/seen", deleteUserSeen)

			movieNum.POST("/longComment", LongComment) //影评
			movieNum.DELETE("/longComment", deleteLongComment)
			movieNum.PUT("/longComment", UpdateLongComment)

			movieNum.POST("/shortComment", ShortComment) //短评
			movieNum.DELETE("/shortComment", deleteShortComment)
			movieNum.PUT("/shortComment", UpdateShortComment)

			commentArea := movieNum.Group("/commentArea")
			{
				commentArea.POST("/", SetCommentArea) //发表讨论区话题
				commentArea.DELETE("/:areaNum", deleteCommentArea)
				commentArea.PUT("/", UpdateArea)
				commentArea.POST("/:areaNum/like", GiveTopicLike) //给某个话题点赞
				commentArea.DELETE("/:areaNum/like", doNotLike)

				comment := commentArea.Group("/:areaNum")
				{
					comment.POST("/comment", GiveComment) //评论某个话题
					comment.DELETE("/comment", deleteComment)
					comment.PUT("/comment", UpdateComment)
					comment.POST("/comment/like", GiveCommentLike) //给某个评论点赞
					comment.DELETE("/comment/like", doNotLikeComment)
				}
			}
		}
	}

	movieGet := engine.Group("/movieInfo")
	{
		movieGet.GET("/:movieNum", GetAMovieInfo) //获取电影的信息
		movieGet.GET("/:movieNum/shortComment", GetShortComment)
		movieGet.GET("/:movieNum/longComment", GetLongComment)
		movieGet.GET("/:movieNum/commentArea", GetCommentArea) //获取讨论区信息
	}

	administrator := engine.Group("/administrator")
	{
		administrator.Use(middleware.JwtToken)
		administrator.Use(middleware.AdministratorToken)

		administrator.POST("/setNewMovie", NewMovie)
	}

	engine.Run(":8090")
}
