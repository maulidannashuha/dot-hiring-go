package tests

import (
	"dot-hiring-go/models"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetUp() *gin.Engine {
	models.ConnectDatabaseTesting()

	return setUpRouter()
}
