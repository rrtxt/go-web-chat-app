package handlers

import (
	"encoding/json"
	"net/http"
	"projects/web-chat-app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

	user.ID = uuid.New()
	user.Password = string(newPassword)

	user, err = user.InsertToDB()
	if err != nil{
		ctx.AbortWithStatusJSON(400, err.Error())
		return
	} else {
		ctx.JSON(200, gin.H{
			"status" : "Success",
			"message" : "Success create new user",
			"user" : user,
		})
	}
}

func GetUserbyUsername(ctx *gin.Context){
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status": "Failed",
			"message" : "Username field is empty",
		})
	}

	user := repositories.User{}
	if err := json.Unmarshal(data, &user); err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}

	resultUser, err := repositories.GetUserFromDB(user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status" : "Failed",
			"message" : "User not found",
		})
	}

	ctx.JSON(200, gin.H{
		"user" : resultUser,
	})
}