package handlers

import (
	"encoding/json"
	"log"
	"projects/web-chat-app/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddUser(ctx *gin.Context) {
	user := repositories.User{}
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

	userFromDB, _ := repositories.GetUserFromDB(user.Username)
	if userFromDB != nil {
		ctx.AbortWithStatusJSON(409, gin.H{
			"status" : "Failed",
			"message" : "Account already exists",
			"data" : userFromDB,
		})
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
		ctx.JSON(200, gin.H{
			"message" : "Success create new user",
			"user" : user,
		})
	}
}

func GetUserbyUsername(ctx *gin.Context){
	user, err := repositories.GetUserFromDB("Test2")

	if err != nil {
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"user" : *user,
	})
}