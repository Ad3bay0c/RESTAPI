package main

import (
	"github.com/Ad3bay0c/WebTesting/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/person", models.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/", models.GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/person/{id}", models.GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/person/{id}", models.UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/person/{id}", models.DeletePersonEndPoint).Methods("DELETE")

	log.Printf("Server Started at localhost:1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
