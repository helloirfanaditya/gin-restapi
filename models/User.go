package models

import (
	"errors"
	"time"
	"trawlcode/database"
	"trawlcode/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

type UserFind struct {
	ID uint `json:"id"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	ID        uint      `json:"id" bindings:"required"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Credential struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func VerifyPassword(password, HashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	var err error
	var user User

	row := database.Db.Where("email = ?", email).First(&user)

	if row.Error != nil {
		return "", errors.New("Wrong Email")
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("wrong Password")
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return "", errors.New("Cannot generate JWT : " + err.Error())
	}

	return token, err
}

func parseUserId(jwt string) float64 {
	userId := utils.VerifyToken(jwt)

	return userId
}

func getUser(id float64) (User, error) {
	var user User
	find := database.Db.Where("id = ?", id).First(&user)

	if find.Error != nil {
		return user, find.Error
	}
	return user, nil
}

type AuthRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
