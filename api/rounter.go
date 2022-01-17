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
	}

	home := engine.Group("/home")
	{
		movie := home.Group("/info")
		{
			movie.GET("/base", MovieBaseInfo)
			movie.GET("/all", MovieInfo)
		}
	}

	engine.Run()
}
