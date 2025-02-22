package main

import (
	"errors"
	"fmt"
)

// creating function for divison
func division(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := division(10, 0)
	if err != nil {
		e := fmt.Errorf("Error is %v.", err.Error())
		fmt.Println(e)
		return
	}
	fmt.Println(result)
}
