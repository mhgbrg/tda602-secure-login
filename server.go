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

const defaultUsername = "admin"
const defaultPassword = "password"

var users []User

type User struct {
	username       string
	hashedPassword []byte
}

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

	//watch out for race conditions (TOCTOU) and timing attack
	for _, user := range users {

		if username == user.username {

			if checkPW(password, user.hashedPassword) {
				fmt.Fprint(w, "Login successful!!!")
			}
		} else {
			fmt.Fprint(w, "Access denied!!!")
		}
	}
}

func main() {
	users = initDatabase()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	http.ListenAndServe(
		"localhost:8080",
		router,
	)
}

func initDatabase() []User {
	file, err := os.Open("database.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	users := make([]User, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		username := parts[0]
		hashedPassword, err := hex.DecodeString(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		user := User{username, hashedPassword}
		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

// open for timing attack
func checkPW(password string, dbHashedPassword []byte) boolean {
	bytePassword := []byte(password)
	hash := sha1.Sum(bytePassword)

}
