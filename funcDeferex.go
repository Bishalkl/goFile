package main

import (
	"errors"
	"fmt"
	"strings"
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

// Implement a function that return another function that generates fibonacci numbers
func fibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		a, b = b, a+b
		return next
	}
}

// Write a function that takes a string and returns a function that counts the number of vowels in that string.
func totalVowels(string string) func() int {
	// conver the string into lowercase
	s := strings.ToLower(string)

	return func() int {
		vowelCounts := 0
		for _, char := range s {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				vowelCounts++
			}

		}
		return vowelCounts
	}
}

// Create a function that checks if a given number is prime or not.
func checkPrime(num int) bool {

	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Write a function that accepts a string and an integer and retgurn the string repeated that many times.
func repString(strings string, num int) string {
	var repstring string
	for i := 0; i < num; i++ {
		repstring += strings
	}
	return repstring
}

func main() {
	x, y := 10, 20
	fmt.Println("Before swapping:", x, y) // Output: Before swapping: 10 20

	swapNum(&x, &y)                      // Passing addresses
	fmt.Println("After swapping:", x, y) // Output: After swapping: 20 10

	fib := fibonacciGenerator() // Create a Fibonacci generator

	// Print the first 10 Fibonacci numbers
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	// Create a vowel counter function for a given string
	countVowels := totalVowels("Hello, GoLang!")

	// Call the function to count vowels
	fmt.Println("Number of vowels:", countVowels()) // Output: 4

	// Test the repeatString function
	fmt.Println(repString("Go", 3))    // Output: GoGoGo
	fmt.Println(repString("Hello", 12)) // Output: HelloHello
}
