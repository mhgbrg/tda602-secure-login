package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// --- MAIN ---

func main() {
	users := loadUsers()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler(users)).Methods("POST")
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}

// --- HANDLERS ---

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

func loginHandler(users map[string][]byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Panic(err)
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		correctPassword, ok := users[username]
		if !ok {
			fmt.Fprintf(w, "Access denied (invalid username)!!!")
			return
		}

		if !checkPassword(password, correctPassword) {
			fmt.Fprintf(w, "Access denied (wrong password)!!!")
			return
		}

		fmt.Fprint(w, "Login successful!!!")
	}
}

// --- HELPERS ---

func loadUsers() map[string][]byte {
	file, err := os.Open("database.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	users := make(map[string][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		username := parts[0]
		password, err := hex.DecodeString(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		users[username] = password
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func checkPassword(password string, correctPassword []byte) bool {
	bytes := []byte(password)
	hash := sha1.Sum(bytes)

	n := len(correctPassword)
	if len(hash) != n {
		return false
	}

	for i := 0; i < n; i++ {
		if hash[i] != correctPassword[i] {
			return false
		}
	}

	return true
}
