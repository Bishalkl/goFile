// Implement an anonymous function inside a for loop to calculate the square of numbers from 1 to 5.
package main

import "fmt"

func main() {
	// creating closer
	square := func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i * i)
		}
	}

	// Invoke the anonymous function
	square()
}
