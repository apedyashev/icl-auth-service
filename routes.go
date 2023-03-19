package main

import (
	"icl-auth/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/register", controllers.RegisterUser)

		api.GET("/user/:id", controllers.GetUserById)
	}
	return router
}
