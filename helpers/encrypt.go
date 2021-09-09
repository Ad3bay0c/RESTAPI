package helpers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", errors.New("Error Hashing Pass")
	}
	return string(newPass), nil
}