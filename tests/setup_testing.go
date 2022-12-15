package tests

import (
	"dot-hiring-go/models"
	"dot-hiring-go/utils"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetUp() *gin.Engine {
	models.ConnectDatabaseTesting()
	utils.SetupRedis()

	return setUpRouter()
}
