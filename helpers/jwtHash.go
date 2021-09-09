package helpers

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
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

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response = map[string]interface{}{
			"message" : "Missing Authorization Token",
		}

		header := r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			json.NewEncoder(w).Encode(response)
			return
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			response["message"] = "Invalid Token"
			json.NewEncoder(w).Encode(response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx := context.WithValue(r.Context(), "userId", claims["userId"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}