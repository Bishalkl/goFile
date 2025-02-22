package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	data string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Occured due to %v", e.data)
}

// creating function for divison
func division(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func main() {
	c, err := division(10, -1)
	if err != nil {
		var d CustomError
		d.data = "Can't assign zero value"
		fmt.Println(d.Error())
		return
	}
	fmt.Println(c)
}
