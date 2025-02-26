// Write an anonymous function that filters out odd numbers from a slice.
package main

import "fmt"

func main() {
	filter := func() func([]int) []int {
		return func(a []int) []int {
			var result []int
			for _, num := range a {
				if num%2 == 0 {
					result = append(result, num)
				}
			}
			return result
		}
	}()

	// Sample slice with odd and even numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Calling the filter function to remove odd numbers
	filteredNumbers := filter(numbers)

	// Print the filtered slice
	fmt.Println(filteredNumbers)
}
