package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	baseURL       = "https://example.com/"
	shortURLChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURLLen   = 6
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var longURL string
	fmt.Print("Enter the long URL: ")
	fmt.Scanln(&longURL)

	shortURL := generateShortURL()
	fmt.Printf("Short URL: %s%s\n", baseURL, shortURL)
}

func generateShortURL() string {
	var sb strings.Builder
	for i := 0; i < shortURLLen; i++ {
		sb.WriteByte(shortURLChars[rand.Intn(len(shortURLChars))])
	}
	return sb.String()
}
