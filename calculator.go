package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num1, num2 float64
    var operator string

    // Prompt the user to enter the first number
    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1)

    // Prompt the user to enter the second number
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2)

    // Prompt the user to enter the operator
    fmt.Print("Enter the operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    var result float64
    var err error

    // Perform the arithmetic operation based on the operator
    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 != 0 {
            result = num1 / num2
        } else {
            fmt.Println("Error: Cannot divide by zero")
            return
        }
    default:
        fmt.Println("Error: Invalid operator")
        return
    }

    // Print the result
    fmt.Println("Result:", result)
}
