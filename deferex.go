package main

import (
	"fmt"
	"log"
	"os"
)

// Write a program that opens a file, writes some text to it, and ensures the file is closed properly using defer.
func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}

	// close after executing
	defer file.Close()

	_, err = file.WriteString("hello world\n")
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("File written successfully")

}
