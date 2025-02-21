package main

import "fmt"

// Function with parameter and return type
func add(a int, b int) int {
	return a + b
}


// Function with multiple return value
func divide( a, b int) (int, int) {
	quotient := a / b
	reminder := a % b

	return quotient, reminder
}


// Function Named Return values
func rectangle(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * ( length + width)
	return 
}

// Function Variadic Function 
func sum(numbers ...int) int {
	total :=0
	for _, num:= range numbers {
		total += num
	}

	return total
} 

func main() {
	// result for add
	result := add(12, 34)
	fmt.Println(result)

	fmt.Println()

	// result for divide
	q, r := divide(12, 6)
	fmt.Println("Quotient:", q, "Remainder:", r)

	fmt.Println()

	// result for rectangle area and perimeter
	a, p := rectangle(12, 34)
	fmt.Println("Area is ", a,  "." , "Perimeter is ", p)

	// result of sum of numbers
	fmt.Println("The total sum is ", sum(1,1,1,1,1,1,1,1,1), ".")

	// defer 
	defer fmt.Println("This will be executed one the program ")
	defer fmt.Println("This will be executed two the program ")
	defer fmt.Println("This will be executed three the program ")
}