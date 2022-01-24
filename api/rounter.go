package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("register", Register)
	engine.POST("login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/change", ChangePassword)
		userGroup.GET("/menu/:username", GetUserInfo) //用户的信息（包括自我介绍
		userGroup.POST("/menu/:username/introduce", SetIntroduce)
		userGroup.GET("/:username/Comment", GetUserComment)    //获取用户的影评和短评
		userGroup.GET("/:username/favorites/wantSee", WantSee) //收藏夹
		userGroup.GET("/:username/favorites/Seen", Seen)
	}

	home := engine.Group("/home")
	{
		home.GET("/research=:find", Find)
		home.GET("/recommend", Recommend)
		home.GET("/:category", FindWithCategory)
	}

	movie := engine.Group("/movie")
	{
		movie.GET("/num=:movieNum", GetAMovieInfo)                    //获取电影的信息
		movie.POST("/num=:movieNum/wantSee", auth, userWantSee)       //用户想看
		movie.POST("/num=:movieNum/seen", auth, userSeen)             //用户看过
		movie.POST("/num=:movieNum/longComment", auth, LongComment)   //影评
		movie.POST("/num=:movieNum/shortComment", auth, ShortComment) //短评
		movie.GET("/GetNum=:movieNum/shortComment", GetShortComment)
		movie.GET("/GetNum=:movieNum/longComment", GetLongComment)
	}

	engine.Run()
}
