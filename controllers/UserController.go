package controllers

import (
	"fmt"
	"trawlcode/models"
	"trawlcode/repositories"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var u []models.User
	users := repositories.GetUser(u)

	c.JSON(200, utils.ResSuccess(200, users))
}

func Find(c *gin.Context) {
	var req models.UserFind

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, "Invalid ID  / ID Must be Integer"))
		return
	}
	getUser, err := repositories.FindUser(req)

	if err != nil {
		c.JSON(400, utils.ResError(400, "Row Not Found"))
		return
	}

	c.JSON(200, utils.ResSuccess(200, getUser))

}

func Update(c *gin.Context) {
	var req models.UserUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Invalid Request : %s", err.Error())))
		return
	}
	updateUser, err := repositories.UpdateUser(req)

	if err != nil {
		c.JSON(400, utils.ResError(400, err.Error()))
		return
	}
	c.JSON(200, utils.ResSuccess(200, updateUser))
}

func Create(c *gin.Context) {
	var req models.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Invalid request : %s", err.Error())))
		return
	}

	if req.Password == "" {
		c.JSON(400, utils.ResError(400, "invalid request Password"))
		return
	}

	var passwordHashed, err = utils.Hash(req.Password)
	if err != nil {
		c.JSON(500, utils.ResError(500, fmt.Sprintf("Hashing password Fail : %s", err.Error())))
		return
	}

	req.Password = passwordHashed

	err = repositories.CreateUser(req)
	if err != nil {
		c.JSON(500, utils.ResError(500, fmt.Sprintf("Hashing password Fail : %s", err.Error())))
		return
	}

	c.JSON(200, utils.ResSuccess(200, map[string]interface{}{
		"message": "success created",
	}))
}

func Delete(c *gin.Context) {
	var req models.UserFind

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("invalid request : %s", err.Error())))
		return
	}

	msg, err := repositories.DeleteUser(req)

	if err != nil {
		c.JSON(400, utils.ResError(400, err.Error()))
		return
	}

	c.JSON(200, utils.ResSuccess(200, msg))

}
