package main

import "fmt"

// Create a function that has a named return value and modify its value inside a defer function.

func modifyval() (value int) {
	// Modify the named return value
	defer func() {
		value++
	}()
	return 12
}

func main() {
	// Call the function and print the result
	result := modifyval()
	fmt.Println("Returned value:", result)
}
