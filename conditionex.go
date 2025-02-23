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

	// var age int

	// // asking the age
	// fmt.Print("Please enter your  age: ")

	// // Check for valid input
	// if _, err := fmt.Scanln(&age); err != nil {
	// 	fmt.Println("Invalid input: Please input valid age.")
	// 	return
	// }

	// // check for valid age
	// if age >= 18 {
	// 	fmt.Println("Yes, you are eligible!")
	// } else {
	// 	fmt.Println("No, you are not eligible!")
	// }

	// var num1, num2, num3 int

	// // ask for input
	// fmt.Print("Please enter the first number: ")

	// // check for vaild input
	// if _, err := fmt.Scanln(&num1); err != nil {
	// 	fmt.Println("Invalid input: Please enter a valid number.")
	// 	return
	// }

	// fmt.Print("Please enter the  second number: ")

	// // check for vaild input
	// if _, err := fmt.Scanln(&num2); err != nil {
	// 	fmt.Println("Invalid input: Please enter a valid number.")
	// 	return
	// }

	// fmt.Print("Please enter the third number: ")

	// // check for vaild input
	// if _, err := fmt.Scanln(&num3); err != nil {
	// 	fmt.Println("Invalid input: Please enter a valid number.")
	// 	return
	// }

	// // check for which is largest
	// if num1 == num2 && num1 == num3 {
	// 	fmt.Println("All numbers are equal.")
	// } else if num1 > num2 && num1 > num3 {
	// 	fmt.Println("The first number is the largest.")
	// } else if num2 > num1 && num2 > num3 {
	// 	fmt.Println("The second number is the largest.")
	// } else {
	// 	fmt.Println("The third number is the largest.")
	// }

}
