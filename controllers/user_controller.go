package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl UserController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get All"})
}