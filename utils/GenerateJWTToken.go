package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id uint) (string, error) {
	life, err := strconv.Atoi("1")

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(life)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("testnet"))
}

func VerifyToken(code string) float64 {
	token, err := jwt.Parse(code, func(token *jwt.Token) (interface{}, error) {
		return []byte("testnet"), nil
	})

	if err != nil {
		return 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0
	}
	fmt.Print(claims)

	id := claims["id"].(float64)
	return id
}
