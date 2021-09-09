package Api_Jwt

import (
	"encoding/json"
	"github.com/Ad3bay0c/WebTesting/db2"
	"github.com/Ad3bay0c/WebTesting/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID			int64			`json:"id,omitempty"`
	Username	string		`json:"username,omitempty"`
	Password	string		`json:"password,omitempty"`
	CreatedAt	time.Time	`json:"created_at,omitempty"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type Profile struct {
	ID			int64		`json:"id,omitempty"`
	Firstname	string	`json:"firstname,omitempty"`
	Lastname	string	`json:"lastname,omitempty"`
}
type Message struct {
	 Message string	`json:"message,omitempty"`
	 Token	string	`json:"token,omitempty"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user = &User{}

	res, err := ioutil.ReadAll(r.Body)
	HandleError(err, w)

	if res == nil {
		json.NewEncoder(w).Encode(Message{Message: "Input Cannot be Empty"})
		return
	}
	_ = json.Unmarshal(res, user)
	//HandleError(err, w)
	if user.Username == "" && user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(Message{Message: "Username and Password is required"})
		return
	}
	if user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(Message{Message: "Username is required"})
		return
	}
	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Password Is Required"})
		return
	}

	if helpers.CheckUser(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Username Already exists"})
		return
	}

	user.ID = time.Now().Unix()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPass, err := helpers.Bcrypt(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: " + err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	user.Password = hashPass

	stmt, err := db2.DB.Prepare("INSERT INTO users (id, username, password) VALUES ($1, $2, $3) ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: " + err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}
	query, err := stmt.Exec(user.ID, user.Username, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: " + err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	id, _ := query.LastInsertId()

	token, err := helpers.CreateToken(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: " + err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	json.NewEncoder(w).Encode(Message{Message: "Registration Successful", Token: token})
}

func HandleError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "An Error Occurred"})
		return
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//var username string

	w.Write([]byte("Login Called"))
}