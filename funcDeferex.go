package main

import "fmt"

// Write a function that takes two numbers and returns their sum, difference, product, and quotient as multiple return values.

func operation(a, b int) (sum int, diff int, product int, quotient int, err error) {
	sum = a + b
	diff = a - b
	product = a * b
	if b != 0 {
		quotient = a / b
	} else {
		err = fmt.Errorf("division by zero is not allowed")
	}
	return
}

func main() {
	a, b := 10, 2
	sum, diff, product, quotient, err := operation(a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
		fmt.Println("Difference:", diff)
		fmt.Println("Product:", product)
		fmt.Println("Quotient:", quotient)
	}
}
