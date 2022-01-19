package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("register", Register)
	engine.POST("login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/change", auth, ChangePassword)
		userGroup.GET("/menu/:username", auth, GetUserInfo)
		userGroup.POST("/menu/:username/introduce", auth, SetIntroduce)
		userGroup.POST("/wantSee", auth, WantSee)
	}

	home := engine.Group("/home")
	{
		home.POST("/find", Find)
		home.GET("recommend", Recommend)
		home.GET("/:category", FindWithCategory)
	}

	movie := engine.Group("/movie")
	{
		movie.GET("/:movieNum", GetAMovieInfo)
		movie.GET("/:movieNum/movieComment", GetMovieComment)
	}

	engine.Run()
}
