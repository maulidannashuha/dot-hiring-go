package controllers

import (
	"dot-hiring-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

var cache utils.Cache

func (ctrl HomeController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome..."})
}
