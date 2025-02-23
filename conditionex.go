package main

import "fmt"

func main() {
	var num int
	// ask to input the number
	fmt.Print("Enter a number: ")

	// check the value is correct or not
	if _, err := fmt.Scanln(&num); err != nil {
		fmt.Println("Invalid input: Please enter an integer")
		return
	}

	// Check if even, odd, or zero
	if num == 0 {
		fmt.Println("It is zero number")
	} else if num%2 == 0 {
		fmt.Println("It is even number")
	} else {
		fmt.Println("It is odd number")
	}

}
