package handlers

import (
	"projects/web-chat-app/models"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context){
	// _, err := ctx.GetRawData()
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(400, "User is not defined")
	// 	return 
	// }

	var userRequest models.UserRequest

	if err := ctx.BindJSON(&userRequest); err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
		return
	}

	// var user models.User

	ctx.JSON(200, gin.H{
		"status" : "Success",
		"Message" : "Success Login",
		"data" : userRequest,
	})
}