package handlers

import "github.com/gin-gonic/gin"

func UserLogin(ctx *gin.Context){
	_, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return 
	}
}