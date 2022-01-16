package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("register", Register)
	engine.POST("login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/change", auth, ChangePassword)
	}

	engine.Run()
}
