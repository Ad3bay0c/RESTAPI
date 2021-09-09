package helpers

import (
	"github.com/Ad3bay0c/WebTesting/db2"
)

func CheckUser(username string) bool {
	var name string
	row := db2.DB.QueryRow("SELECT username FROM users WHERE username = $1", username)
	_ = row.Scan(&name)
	if name == "" {
		return false
	} else {
		return true
	}
}
