package controllers

import (
	"fmt"
	"trawlcode/models"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Index(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(200, utils.ResSuccess(200, users))
}

func Find(c *gin.Context) {
	var req models.UserFind
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, "Invalid ID  / ID Must be Integer"))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", req.ID).First(&user).Error; err != nil {
		c.JSON(400, utils.ResError(400, "User Not Found !"))
		return
	}

	c.JSON(200, utils.ResSuccess(200, user))

}

func Update(c *gin.Context) {
	var req models.UserUpdate
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Invalid Request : %s", err.Error())))
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", req.ID).First(&user).Error; err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("Failed to Update Because : %s", err.Error())))
		return
	}

	user.Name = req.Name
	user.Email = req.Email
	if req.Password != "" {
		var passwordHashed, err = utils.Hash(req.Password)
		if err != nil {
			c.JSON(500, utils.ResError(500, fmt.Sprintf("Hashing password Fail : %s", err.Error())))
		}
		user.Password = passwordHashed
	}

	db.Save(&user)

	c.JSON(200, utils.ResSuccess(200, user))
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

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHashed,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(200, utils.ResSuccess(200, "success created"))
}

func Delete(c *gin.Context) {
	var req models.UserFind
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResError(400, fmt.Sprintf("invalid request : %s", err.Error())))
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", req.ID).First(&user).Error; err != nil {
		c.JSON(400, utils.ResError(400, "User Not Found"))
		return
	}
	db.Delete(&user)

	c.JSON(200, utils.ResSuccess(200, "Deleted !"))

}
