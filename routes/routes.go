package routes

import (
	"projects/web-chat-app/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" : "Heyy",
		})
	})
	router.POST("/user", handlers.AddUser)
}