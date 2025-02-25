package main

import (
	"fmt"
	"time"
)

// Implement a function that tracks the execution time of another function using defer and time.Now().

// first create function calculate it
func executionTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}

// create a main function
func main() {
	defer executionTime(time.Now())

	// simulate time processing
	time.Sleep(2 * time.Second)

	fmt.Println("Function execution complete")
}
