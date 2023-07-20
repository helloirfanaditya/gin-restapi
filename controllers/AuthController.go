package controllers

import (
	"fmt"
	"trawlcode/models"
	"trawlcode/repositories"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Invalid Request : %s", err.Error())))
		return
	}

	token, err := repositories.AuthenticationLogin(req)
	if err != nil {
		c.JSON(500, utils.ResError(500, err.Error()))
		return
	}
	c.JSON(200, utils.ResSuccess(200, map[string]interface{}{
		"token": token,
	}))
}

func Register(c *gin.Context) {
	c.JSON(200, utils.ResSuccess(200, map[string]interface{}{
		"message": "ok",
	}))
}
