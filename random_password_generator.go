package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var length int
	fmt.Print("Enter the length of the password: ")
	fmt.Scanln(&length)

	var useDigits, useSymbols bool
	var symbolChoice string

	fmt.Print("Include digits (0-9)? (y/n): ")
	fmt.Scanln(&useDigits)
	fmt.Print("Include symbols (!@#$%^&*()_+)? (y/n): ")
	fmt.Scanln(&symbolChoice)

	if symbolChoice == "y" {
		useSymbols = true
	} else {
		useSymbols = false
	}

	password := generatePassword(length, useDigits, useSymbols)
	fmt.Println("Generated password:", password)
}

func generatePassword(length int, useDigits, useSymbols bool) string {
	var charset string
	if useDigits {
		charset += "0123456789"
	}
	if useSymbols {
		charset += "!@#$%^&*()_+"
	}
	charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
