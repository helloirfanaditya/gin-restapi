package repositories

import (
	"errors"
	"fmt"
	"trawlcode/database"
	"trawlcode/models"
	"trawlcode/utils"
)

func GetUser(payload []models.User) interface{} {
	database.Db.Find(&payload)
	return payload
}

func CreateUser(payload models.UserCreate) error {
	user := models.User(payload)
	database.Db.Create(&user)

	return nil
}

func FindUser(payload models.UserFind) (*models.User, error) {
	var user models.User
	r := database.Db.Where("id = ?", payload.ID).First(&user)

	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

func UpdateUser(payload models.UserUpdate) (*models.User, error) {
	var user models.User
	if err := database.Db.Where("id = ?", payload.ID).First(&user).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to update Because : %s", err.Error()))
	}

	user.Name = payload.Name
	user.Email = payload.Email
	if payload.Password != "" {
		var passwordHashed, err = utils.Hash(payload.Password)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Hashing password Fail : %s", err.Error()))
		}
		user.Password = passwordHashed
	}

	database.Db.Save(&user)

	return &user, nil
}

func DeleteUser(payload models.UserFind) (string, error) {
	var user models.User

	if err := database.Db.Where("id = ?", payload.ID).First(&user).Error; err != nil {
		return "", errors.New("User Not Found")
	}
	database.Db.Delete(&user)

	return "Deleted !", nil
}
