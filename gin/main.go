package main

import (
	"gin-mongo-api/configs" //add this
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}
