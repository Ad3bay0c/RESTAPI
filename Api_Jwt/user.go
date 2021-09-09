package Api_Jwt

import (
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message":"Called"})
	//fmt.Fprintf(w, "Get User Called")
}
