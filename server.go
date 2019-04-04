package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const defaultUsername = "admin"
const defaultPassword = "password"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("login.html")
	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Panic(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == defaultUsername && password == defaultPassword {
		fmt.Fprint(w, "Login successful!!!")
	} else {
		fmt.Fprint(w, "Access denied!!!")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}
