package main

import (
	"fmt"
	"github.com/Ad3bay0c/WebTesting/models"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v: %T", r.Method, r.Method)

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println("ID", r.Form["id"])
	fmt.Println("Username", r.Form["user"])
	fmt.Println("Method", r.Method)

	fmt.Fprintf(w,"Started in peace")
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", models.LoginUser)

	fmt.Println("Server Started at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
