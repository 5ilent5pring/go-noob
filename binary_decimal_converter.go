package main

import (
	"fmt"
	"strconv"
)

func main() {
	var choice int
	fmt.Println("Choose conversion:")
	fmt.Println("1. Decimal to Binary")
	fmt.Println("2. Binary to Decimal")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var decimal int
		fmt.Print("Enter decimal number: ")
		fmt.Scanln(&decimal)
		binary := decimalToBinary(decimal)
		fmt.Printf("Binary representation: %s\n", binary)
	case 2:
		var binary string
		fmt.Print("Enter binary number: ")
		fmt.Scanln(&binary)
		decimal, err := binaryToDecimal(binary)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Decimal representation: %d\n", decimal)
	default:
		fmt.Println("Invalid choice.")
	}
}

func decimalToBinary(decimal int) string {
	return strconv.FormatInt(int64(decimal), 2)
}

func binaryToDecimal(binary string) (int, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}
