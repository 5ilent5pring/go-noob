package main

import (
    "fmt"
    "strings"
)

func main() {
    // Prompt the user to enter a string
    fmt.Print("Enter a string: ")

    var input string
    fmt.Scanln(&input)

    // Remove spaces from the input string
    input = strings.ReplaceAll(input, " ", "")

    // Convert the input string to lowercase
    input = strings.ToLower(input)

    // Check if the input string is a palindrome
    if isPalindrome(input) {
        fmt.Println("Yes, it is a palindrome.")
    } else {
        fmt.Println("No, it is not a palindrome.")
    }
}

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
    length := len(s)
    for i := 0; i < length/2; i++ {
        if s[i] != s[length-i-1] {
            return false
        }
    }
    return true
}
