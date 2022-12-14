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

	r.GET("/users", userController.GetAll)
	r.POST("/users", userController.Store)
	r.PUT("/users/:id", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	r.GET("/books", bookController.GetAll)
	r.POST("/books", bookController.Store)
	r.PUT("/books/:id", bookController.Update)
	r.DELETE("/books/:id", bookController.Delete)

	return r
}
