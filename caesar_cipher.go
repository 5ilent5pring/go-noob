package main

import (
	"fmt"
)

func main() {
	var choice int
	fmt.Println("Choose operation:")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var plaintext string
		var shift int
		fmt.Print("Enter plaintext: ")
		fmt.Scanln(&plaintext)
		fmt.Print("Enter shift (positive integer): ")
		fmt.Scanln(&shift)
		ciphertext := encrypt(plaintext, shift)
		fmt.Println("Ciphertext:", ciphertext)
	case 2:
		var ciphertext string
		var shift int
		fmt.Print("Enter ciphertext: ")
		fmt.Scanln(&ciphertext)
		fmt.Print("Enter shift (positive integer): ")
		fmt.Scanln(&shift)
		plaintext := decrypt(ciphertext, shift)
		fmt.Println("Plaintext:", plaintext)
	default:
		fmt.Println("Invalid choice.")
	}
}

func encrypt(plaintext string, shift int) string {
	var result string
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}

func decrypt(ciphertext string, shift int) string {
	return encrypt(ciphertext, -shift)
}
