package controllers

import (
	"trawlcode/models"
	"trawlcode/repositories"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
)

func IndexPosts(c *gin.Context) {
	var u []models.Post

	posts := repositories.ListPost(u)

	c.JSON(200, utils.ResSuccess(200, posts))
}
