package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Prompt the user to enter a message
    fmt.Print("Enter a message: ")

    // Create a reader to read user input from the console
    reader := bufio.NewReader(os.Stdin)

    // Read the user input
    userInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Print the user input back to the console
    fmt.Println("You entered:", userInput)
}
