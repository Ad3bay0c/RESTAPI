package Api_Jwt

import (
	"encoding/json"
	"github.com/Ad3bay0c/WebTesting/db2"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userId")

	var user User
	row := db2.DB.QueryRow("SELECT id, username FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
	}

	json.NewEncoder(w).Encode(user)
}
