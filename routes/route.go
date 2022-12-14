package routes

import (
	"dot-hiring-go/controllers"

	"github.com/gin-gonic/gin"
)

var homeController = controllers.HomeController{}
var userController = controllers.UserController{}
var bookController = controllers.BookController{}

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", homeController.Welcome)

	r.GET("/user", userController.GetAll)
	r.POST("/user", userController.Store)
	r.PUT("/user/:userId", userController.Update)
	r.DELETE("/user/:userId", userController.Delete)

	r.GET("/user/:userId/books", bookController.GetAll)
	r.POST("/user/:userId/books", bookController.Store)
	r.PUT("/user/:userId/books/:id", bookController.Update)
	r.DELETE("/user/:userId/books/:id", bookController.Delete)

	return r
}
