package main

import (
	"dot-hiring-go/models"
	"dot-hiring-go/routes"
	"dot-hiring-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	utils.SetupRedis()

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r = routes.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
