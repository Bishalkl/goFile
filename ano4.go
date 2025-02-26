// Write a function that returns a closure, which takes an integer and multiplies it by a given factor.
package main

import (
	"fmt"
)

func factor(factor int) func(int) int {
	return func(multiplier int) int {
		return factor * multiplier
	}
}

func main() {
	// Creating a closure with a factor of 5
	multipart := factor(5)

	fmt.Println(multipart(2))

}
