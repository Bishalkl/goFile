// Create a function that returns a closure which keeps track of the sum of numbers added to it.

package main

import "fmt"

func sumTrack() func(int) int {
	var sum int
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	// create a new sum closure
	add := sumTrack()

	// Add numbers and print the updated sum each time
	fmt.Println(add(5))  // Output: 5
	fmt.Println(add(10)) // Output: 15
	fmt.Println(add(3))  // Output: 18
	fmt.Println(add(7))  // Output: 25
}
