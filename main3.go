package main

import (
	"github.com/Ad3bay0c/WebTesting/Api_Jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", Api_Jwt.LoginHandler).Methods("POST")
	router.HandleFunc("/register", Api_Jwt.RegisterHandler).Methods("POST")
	router.HandleFunc("/user", Api_Jwt.GetUser).Methods("GET")

	log.Println("Server Started at Localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", router))

}
