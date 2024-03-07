package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Prompt the user to enter the filename
    fmt.Print("Enter the filename: ")
    var filename string
    fmt.Scanln(&filename)

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Create a scanner to read from the file
    scanner := bufio.NewScanner(file)

    // Initialize word count variable
    wordCount := 0

    // Scan through the file line by line
    for scanner.Scan() {
        // Split each line into words
        words := strings.Fields(scanner.Text())
        // Increment the word count by the number of words in the line
        wordCount += len(words)
    }

    // Check for any errors encountered during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the word count
    fmt.Println("Word count:", wordCount)
}
