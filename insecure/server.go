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
	loginTemplate, err := template.ParseFiles("login.html")
	if err != nil {
		log.Fatal(err)
	}
	logoutTemplate, err := template.ParseFiles("logout.html")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler(loginTemplate, logoutTemplate)).Methods("GET")
	router.HandleFunc("/login", loginHandler(users)).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	httpHost := os.Getenv("HOSTNAME") + ":" + os.Getenv("HTTP_PORT")
	log.Printf("listening for http requests on %s", httpHost)
	log.Fatal(http.ListenAndServe(httpHost, router))
}

// --- HANDLERS ---

func indexHandler(loginTemplate, logoutTemplate *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err == http.ErrNoCookie {
			err = loginTemplate.Execute(w, nil)
		} else {
			err = logoutTemplate.Execute(w, cookie.Value)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loginHandler(users map[string][]byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		username := strings.TrimSpace(r.Form.Get("username"))
		password := r.Form.Get("password")

		correctPassword, ok := users[username]
		if !ok {
			fmt.Fprintf(w, "Access denied!!!")
			return
		}

		if !checkPassword(password, correctPassword) {
			fmt.Fprintf(w, "Access denied!!!")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: username,
		})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusFound)
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
