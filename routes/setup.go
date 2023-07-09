package routes

import (
	UserController "trawlcode/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Route User
	r.GET("/users", UserController.Index)
	r.GET("/user", UserController.Find)
	r.POST("/user/update", UserController.Update)
	r.POST("/user/create", UserController.Create)
	r.POST("/user/delete", UserController.Delete)

	return r
}
