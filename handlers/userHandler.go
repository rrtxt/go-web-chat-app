package handlers

import (
	"encoding/json"
	"projects/web-chat-app/handlers/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct{
	Username string `json:"username"`
	Password string `json:"password"`
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
		ctx.AbortWithStatusJSON(400, err.Error())
		return 
	}
	
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		ctx.AbortWithStatusJSON(400, err.Error())
		return 
	}

	user.Password = string(newPassword)

	user, err = user.InsertToDB()
	if err != nil{
		ctx.AbortWithStatusJSON(400, err.Error())
		return
	} else {
		// userResponse := UserResponse{
		// 	Username: user.Username,
		// 	Password: user.Password,
		// }
		ctx.JSON(200, gin.H{
			"message" : "Success create new user",
			"user" : user,
		})
	}
}