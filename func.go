package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello world"
	printMe(printValue)

	numerator := 11
	denominator := 0

	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
