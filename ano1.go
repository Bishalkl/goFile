// Implement an anonymous function that accepts two integers and returns their sum, then assign it to a variable and call it.

package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(3, 4))

}
