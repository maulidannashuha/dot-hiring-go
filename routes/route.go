package routes

import (
	"dot-hiring-go/controllers"

	"github.com/gin-gonic/gin"
)

var homeController = controllers.HomeController{}
var userController = controllers.UserController{}

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", homeController.Welcome)

	r.GET("/users", userController.GetAll)

	return r
}