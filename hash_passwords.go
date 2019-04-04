package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	username string
	password string
}

func main() {
	users := readUsers()
	for _, user := range users {
		bytePassword := []byte(user.password)
		hash := sha1.Sum(bytePassword)
		fmt.Printf("%s\t%x\n", user.username, hash)
	}
}

func readUsers() []User {
	file, err := os.Open("users.txt")
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
		password := parts[1]
		user := User{username, password}
		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
