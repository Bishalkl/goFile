package main

import "fmt"

func mergedSortedArrays(arr1, arr2 []int) []int {
	// create merged array to sort
	merged := make([]int, 0, len(arr1)+len(arr2))

	// intialized pointer for array
	i, j := 0, 0

	// now loop through it and merged
	for i < len(arr1) && j < len(arr2) {

		// condtion
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append any remaining elements from arr1
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	// Append any remaining elements fro arr2
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged

}

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}
	merged := mergedSortedArrays(arr1, arr2)

	// output
	fmt.Println("Merged Sorted Array: ", merged)
}
