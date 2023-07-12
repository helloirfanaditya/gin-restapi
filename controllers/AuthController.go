package controllers

import (
	"fmt"
	"trawlcode/models"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identity = "id"

func Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Invalid Request : %s", err.Error())))
		return
	}

	token, err := models.LoginCheck(req.Email, req.Password)

	if err != nil {
		c.JSON(500, utils.ResError(500, fmt.Sprintf("error : %s", err.Error())))
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
