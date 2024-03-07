package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const maxAttempts = 5
	target := rand.Intn(100) + 1 // Generate a random number between 1 and 100

	fmt.Println("Welcome to the Guess the Number game!")
	fmt.Println("I've picked a number between 1 and 100. You have", maxAttempts, "attempts to guess it.")

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		var guess int
		fmt.Printf("Attempt %d: Enter your guess: ", attempts)
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the number correctly!")
			return
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}

	fmt.Println("Sorry, you've run out of attempts. The number was", target)
}
