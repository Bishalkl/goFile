package main

import "fmt"

//Create a function where multiple defer statements are used. What will be the order of execution?

func main() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")
	fmt.Println("end")

	// output we will be start and end and 5 4 3 2 1 beacause defer function execute after function execution and it will be act like lifo like stack
}
