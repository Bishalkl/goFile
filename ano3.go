// Create a closure that keeps track of how many times a function is called and return the count
package main

import "fmt"

func main() {
	// creating closures
	trackCounts := func() func() int {
		var count int = 0
		return func() int {
			count++
			return count
		}
	}()

	// calling manytimes function
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())
	fmt.Println(trackCounts())

}
