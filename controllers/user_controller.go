package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl UserController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Get All"})
	}
}

func Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get All"})
}
