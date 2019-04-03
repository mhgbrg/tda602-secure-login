package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}
