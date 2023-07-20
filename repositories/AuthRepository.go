package repositories

import (
	"trawlcode/models"
)

func AuthenticationLogin(payload models.AuthRequest) (string, error) {
	token, err := models.LoginCheck(payload.Email, payload.Password)

	if err != nil {
		return "", err
	}

	return token, nil
}
