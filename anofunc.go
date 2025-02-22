package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, go")


	// Anonymous function invoked immediately
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	// Anonymous function assigned to a variable
	greet := func(name string) {
		fmt.Println("Hello, ", name)
	}

	greet("Bishal")

	// Anonymous function with return value

	add := func(a , b int) int {
		return a + b
	} 

	result := add(12, 34)
	fmt.Println("Sum:", result)

}