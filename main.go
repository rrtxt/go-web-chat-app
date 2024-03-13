package main

import (
	"log"
	"projects/web-chat-app/database"
	"projects/web-chat-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:" , err)
	}

	router := gin.Default()
	address := "localhost:8000"
	routes.InitRoutes(router)

	database.ConnectDB()

	router.Run(address)
}