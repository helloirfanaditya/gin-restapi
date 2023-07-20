package repositories

import (
	"errors"
	"fmt"
	"trawlcode/database"
	"trawlcode/models"
	"trawlcode/utils"
)

type Result struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func GetUser(payload []models.User) interface{} {
	var result []Result

	database.Db.Find(&payload)

	for _, data := range payload {
		r := Result{
			ID:        data.ID,
			Name:      data.Name,
			Email:     data.Email,
			CreatedAt: data.CreatedAt.Format("January 2, 2006"),
			UpdatedAt: data.UpdatedAt.Format("January 2, 2006"),
		}
		result = append(result, r)
	}
	if len(result) == 0 {
		return []string{}
	}
	return result
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
