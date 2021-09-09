package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func init () {
	os.Setenv("SECRET_KEY", "mySecreteJWTKey")
}
func CreateToken(userId int64) (string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Authorized": true,
		"expiry": time.Now().Add(time.Minute * 30).Unix(),
		"issuedAt": time.Now().Unix(),
		"userId": userId,
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}