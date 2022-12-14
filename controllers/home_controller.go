package controllers

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (ctrl HomeController) Welcome(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "Welcome..."})
}