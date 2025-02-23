package main

import "fmt"

func main() {
	// var num int
	// // ask to input the number
	// fmt.Print("Enter a number: ")

	// // check the value is correct or not
	// if _, err := fmt.Scanln(&num); err != nil {
	// 	fmt.Println("Invalid input: Please enter an integer")
	// 	return
	// }

	// // Check if even, odd, or zero
	// if num == 0 {
	// 	fmt.Println("It is zero number")
	// } else if num%2 == 0 {
	// 	fmt.Println("It is even number")
	// } else {
	// 	fmt.Println("It is odd number")
	// }

	// var num int

	// // first ask the number
	// fmt.Print("Please enter a number: ")

	// // check it is correct input or not
	// if _, err := fmt.Scanln(&num); err != nil {
	// 	fmt.Println("Invalid input: Please enter an integer.")
	// 	return
	// }

	// // main checking
	// if num == 0 {
	// 	fmt.Println("It is a zero number.")
	// } else if num > 0 {
	// 	fmt.Println("it is a positive  number.")
	// } else {
	// 	fmt.Println("it is a negative number.")
	// }

	var age int

	// asking the age 
	fmt.Print("Please enter your  age: ")

	// Check for valid input
	if _, err := fmt.Scanln(&age); err != nil {
		fmt.Println("Invalid input: Please input valid age.")
		return 
	}

	// check for valid age 
	if age >= 18 {
		fmt.Println("Yes, you are eligible!")
	} else {
		fmt.Println("No, you are not eligible!")
	}



}
