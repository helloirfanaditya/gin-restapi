package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	var pwd = []byte(password)

	Hashed, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	return string(Hashed), err
}
