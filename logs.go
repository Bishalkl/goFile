package main

import (
	"fmt"
)

func logExecution() {

	// Log the start of execution
	fmt.Println("Start of execution")

	// Defer the log for the end of execution
	defer fmt.Println("End of execution")

	// Simulate some work
	fmt.Println("Executing some tasks...")
}

func main() {
	// Call the logExecution function
	logExecution()
}
