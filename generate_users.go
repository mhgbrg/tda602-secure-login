package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var passwordChars = []rune(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")

func main() {
	rand.Seed(time.Now().UnixNano())

	usernames := readFile("names.txt")
	words := readFile("words.txt")

	for i := 0; i < 50; i++ {
		username := usernames[i]
		word := words[rand.Intn(len(words))]
		password := word + strconv.Itoa(rand.Intn(10))
		fmt.Printf("%s\t%s\n", username, password)
	}

	for i := 50; i < 75; i++ {
		username := usernames[i]
		password := randomPassword(7)
		fmt.Printf("%s\t%s\n", username, password)
	}

	for i := 75; i < 100; i++ {
		username := usernames[i]
		password := randomPassword(25)
		fmt.Printf("%s\t%s\n", username, password)
	}
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func randomPassword(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = passwordChars[rand.Intn(len(passwordChars))]
	}
	return string(b)
}
