package main

import "fmt"

func main() {

	// Normal if condition in go lang 
	x:=10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {
		fmt.Println("x is less than or equal to 5")
	} else {
		fmt.Println("x is equal to 5")
	}

	fmt.Println()

	// short Variable Declaration in if statement
	if num:=7; num%2 == 0 {
		fmt.Println("Even number")
	} else if num == 0{
		fmt.Println("Equal to zero")
	} else {
		fmt.Println("Odd number")
	}

	fmt.Println()

	// Normal Switch Statement and multiple cases 
	day := "Thursday"
	
	switch day {
	case "Monday":
		fmt.Println("start of the workweek")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Workweek is going")
	case "Friday":
		fmt.Println("Weekend is near")
	default:
		fmt.Println("Regular day")
	} 

	fmt.Println()

	// Switch Without an Expression
	num2:= 20

	switch {
	case num2 < 20:
		fmt.Println("Greater than 20")
	case num2 <= 20 && num2 >= 10:
		fmt.Println("Num is smaller than 20 and greater 10")  
	default:
		fmt.Println("Smaller than 10")
	} 

	fmt.Println()

	// Ternary Operator Alternative in go 
	num3:=7
	result  := "Even"
	if num3% 2 != 0 {
		result = "odd"
	}
	fmt.Println(result)

}
