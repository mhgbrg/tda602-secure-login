package main

import (
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}
