package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter the directory path
	fmt.Print("Enter the directory path: ")
	var dirPath string
	fmt.Scanln(&dirPath)

	// Read the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate over files in the directory
	for _, fileInfo := range files {
		// Check if the file is a regular file and has a .txt extension
		if fileInfo.Mode().IsRegular() && strings.HasSuffix(fileInfo.Name(), ".txt") {
			// Open the file
			filePath := dirPath + "/" + fileInfo.Name()
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				continue
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
				continue
			}

			// Print the word count for the current file
			fmt.Printf("Word count for %s: %d\n", fileInfo.Name(), wordCount)
		}
	}
}
