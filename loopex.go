package main

import (
	"fmt"
)

func dosomethingwithrune(r rune) {
	fmt.Println(string(r))
}

func main() {
	// //print number from 1 to 20 using a for loop
	// for i := 1; i >= 20; i++ {
	// 	fmt.Print(i, " ")
	// }

	// fmt.Println()

	// //print number from 10 to 1 using a for loop
	// for i := 10; i >= 1; i-- {
	// 	fmt.Print(i, " ")
	// }

	// fmt.Println()

	// // print the sum of numbers from 1 to 100
	// var sum int = 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum)

	// var i int
	// for {
	// 	fmt.Print("Enter a number (0 to exit): ")
	// 	if _, err := fmt.Scanln(&i); err != nil {
	// 		fmt.Println("Invalid input. Please enter an integer")
	// 		continue
	// 	}

	// 	if i == 0 {
	// 		fmt.Println("Exiting program...")
	// 		break
	// 	}
	// }

	// var num int
	// fmt.Printf("Enter a number: ")
	// if _, err := fmt.Scanln(&num); err != nil {
	// 	fmt.Println("Invalid input: Please enter an integer")
	// 	return
	// }
	// // now loop for multiplication
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%d * %d = %d\n", num, i, num*i)
	// }

	// var slice1 = []int{3, 7, 2, 9, 5}
	// var sum int = 0
	// for _, r := range slice1 {
	// 	sum += r
	// }
	// fmt.Printf("Sum: %d", sum)

	// myMap := map[string]int{"a": 10, "b": 20, "c": 30}
	// for key, value := range myMap {
	// 	fmt.Printf("Key:  %s, Value: %d\n", key, value)
	// }

	// // Print a 5*5 square of * using loops
	// for i := 0; i < 5; i++ { //for outerloop
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Print("* ")
	// 	}
	// 	fmt.Println()
	// }

	// Print a right-angled triangle of numbers
	//for outer loop
	// for i := 1; i <= 5; i++ {
	// 	// for inner loop
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d ", j)
	// 	}
	// 	fmt.Println()
	// }

	// FizzBuzz challenge
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}

}
