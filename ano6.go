// Create a closure that maintains a counter and allows incrementing or resetting the counter.

package main

import "fmt"

func main() {
	counter := func() func(string) int {
		count := 0
		return func(action string) int {
			if action == "increment" {
				count++
			} else if action == "reset" {
				count = 0
			}
			return count
		}
	}()

	// Testing the closure
	fmt.Println("Count after increment:", counter("increment")) // Output: 1
	fmt.Println("Count after increment:", counter("increment")) // Output: 2
	fmt.Println("Count after reset:", counter("reset"))         // Output: 0
	fmt.Println("Count after increment:", counter("increment")) // Output: 1
}
