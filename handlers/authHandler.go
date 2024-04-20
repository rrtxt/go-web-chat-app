package handlers

import (
	"log"
	"projects/web-chat-app/models"
	"projects/web-chat-app/repositories"
	"projects/web-chat-app/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	user, err := repositories.GetUserFromDB(userRequest.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status" : "Failed",
			"message" : "Error on fetching data",
		})
		return 
	}

	if user == nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"status" : "Failed",
			"message" : "User not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		resPwd, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil{
			log.Println(err.Error())
		}
		ctx.AbortWithStatusJSON(401, gin.H{
			"status" : "Failed",
			"message" : "Wrong Password",
			"user" : (user.Password),
			"userReq" : resPwd,
		})
		return
	}

	claims := utils.JWTClaims{Id: user.ID.String(), Name: user.Username}
	token, err := claims.Encode()
	
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status" : "Failed",
			"message" : "Something went wrong!",
		})
	}

	ctx.JSON(200, gin.H{
		"status" : "Success",
		"message" : "Success Login",
		"data" : user,
		"token" : token,
	})
}