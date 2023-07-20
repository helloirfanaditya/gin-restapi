package routes

import (
	AuthController "trawlcode/controllers"
	PostController "trawlcode/controllers"
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
	r.GET("/users", UserController.Index)
	authenticationRoute.GET("/user", UserController.Find)
	authenticationRoute.POST("/user/update", UserController.Update)
	r.POST("/user/create", UserController.Create)
	authenticationRoute.POST("/user/delete", UserController.Delete)

	// Post Route
	authenticationRoute.GET("/posts", PostController.IndexPosts)
	return r
}
