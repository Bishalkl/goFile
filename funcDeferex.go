package main

import (
	"errors"
	"fmt"
)

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

func factorial(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Factorial is not defined for negative numbers")
	}

	if num == 0 || num == 1 {
		return 1, nil
	}

	// Recursive call
	fact, err := factorial(num - 1)
	if err != nil {
		return 0, err
	}

	return num * fact, nil

}

// write a function that accepts a variable number of integer arguments (variadic function) and return their sum.

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

// Write a function that swaps two integer using pointers in Go.
func swapNum(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 20
	fmt.Println("Before swapping:", x, y) // Output: Before swapping: 10 20

	swapNum(&x, &y)                      // Passing addresses
	fmt.Println("After swapping:", x, y) // Output: After swapping: 20 10
}
