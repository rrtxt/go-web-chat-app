package handlers

import (
	"encoding/json"
	"projects/web-chat-app/handlers/repository"

	"github.com/gin-gonic/gin"
)

type UserResponse struct{
	Username string `json:"username"`
}

func AddUser(ctx *gin.Context) {
	user := repository.User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return 
	}

	user, err = user.InsertToDB()
	if err != nil{
		ctx.AbortWithStatusJSON(400, "Failed to create new user")
		return
	} else {
		userResponse := UserResponse{
			Username: user.Username,
		}
		ctx.JSON(200, gin.H{
			"message" : "Success create new user",
			"user" : userResponse,
		})
	}
}