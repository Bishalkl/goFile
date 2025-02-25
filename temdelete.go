package main

import (
	"fmt"
	"log"
	"os"
)

// Function to delete the file safely
func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Println("Warning: Error deleting file:", err) // Log error instead of terminating
		return
	}
	fmt.Println("Successfully deleted", filename)
}

func main() {
	// Creating a temporary file in the system's temp directory
	file, err := os.CreateTemp("", "example") // Generates a unique temp file
	if err != nil {
		log.Fatal("Error creating file:", err)
	}

	// Ensure the file is closed before deletion
	defer file.Close()
	defer deleteFile(file.Name())

	fmt.Println("Temporary file created:", file.Name())
}
