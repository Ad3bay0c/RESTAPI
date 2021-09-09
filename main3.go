package main

import (
	"github.com/Ad3bay0c/WebTesting/Api_Jwt"
	"github.com/Ad3bay0c/WebTesting/helpers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", Api_Jwt.LoginHandler).Methods("POST")
	router.HandleFunc("/register", Api_Jwt.RegisterHandler).Methods("POST")

	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.Use(helpers.IsAuthorized)

	subRouter.HandleFunc("/user", Api_Jwt.GetUser).Methods("GET")

	sR := router.PathPrefix("/user").Subrouter()
	sR.Use(helpers.IsAuthorized)

	sR.HandleFunc("/contacts", Api_Jwt.GetAllContacts).Methods("GET")
	sR.HandleFunc("/contact/{id}", Api_Jwt.DeleteContact).Methods("DELETE")
	sR.HandleFunc("/contact/{id}", Api_Jwt.UpdateContact).Methods("PUT")
	sR.HandleFunc("/contact", Api_Jwt.CreateContact).Methods("POST")
	sR.HandleFunc("/contact/{id}",Api_Jwt.GetContact).Methods("GET")

	log.Println("Server Started at Localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", router))

}
