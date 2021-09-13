package helpers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", errors.New("error Hashing Pass")
	}
	return string(newPass), nil
}

func VerifyPassword(pass string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {

		return false
	}
	return true
}