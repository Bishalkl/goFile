package main

import "fmt"

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

	var i int
	for {
		fmt.Print("Enter a number (0 to exit): ")
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println("Invalid input. Please enter an integer")
			continue
		}

		if i == 0 {
			fmt.Println("Exiting program...")
			break
		}
	}

}
