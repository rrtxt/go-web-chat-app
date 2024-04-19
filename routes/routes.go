package routes

import (
	"projects/web-chat-app/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/user", handlers.GetUserbyUsername)
	router.POST("/user", handlers.AddUser)
	router.POST("/login", handlers.UserLogin)
}