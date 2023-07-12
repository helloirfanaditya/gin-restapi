package routes

import (
	AuthController "trawlcode/controllers"
	UserController "trawlcode/controllers"
	"trawlcode/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	// Authentication
	authenticationRoute := r.Group("", middlewares.Logger())

	r.POST("/login", AuthController.Login)
	r.POST("/register", AuthController.Register)

	// Route User
	authenticationRoute.GET("/users", UserController.Index)
	authenticationRoute.GET("/user", UserController.Find)
	authenticationRoute.POST("/user/update", UserController.Update)
	authenticationRoute.POST("/user/create", UserController.Create)
	authenticationRoute.POST("/user/delete", UserController.Delete)

	return r
}
