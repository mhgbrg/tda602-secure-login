package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprint(w, "Hej")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}
