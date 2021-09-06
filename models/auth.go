package models

import (
	md52 "crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type LoginDetails struct {
	Username	string
	Password	string
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v: %T", r.Method, r.Method)
	t := time.Now().Unix()
	h := md52.New()
	io.WriteString(h, strconv.FormatInt(t, 10))

	token := fmt.Sprintf("%x", h.Sum(nil))

	if r.Method == "GET" {
		t, _ := template.ParseFiles("frontend/login.gtpl")
		t.Execute(w, token)
	} else if r.Method == "POST" {
		r.ParseForm()

		if r.Form.Get("token") != "" && r.Form.Get("token") == token {
			fmt.Fprintf(w, "Username:\t%v\nPassword:\t%v", r.Form["username"][0], r.FormValue("password"))

		} else {
			http.Redirect(w, r, "/login", http.StatusOK)
		}
		//l := LoginDetails{
		//	Username: r.Form["username"][0],
		//	Password: r.Form["password"][0],
		//}
		//ParseToFile("frontend/login.gtpl", l, w)
	} else {
		t, _ := template.ParseFiles("frontend/login.gtpl")
		t.Execute(w, nil)

	}
}

func ParseToFile(filename string, data interface{}, w http.ResponseWriter) {
	t, _ := template.ParseFiles(filename)

	t.Execute(w, data)
}
