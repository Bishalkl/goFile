package main

import (
	"fmt"
	"slices"
)

func main() {
	var zeroArray [5]int
	var myArray = [3]int{5, 10, 15}
	var cars = [5]string{1: "Volve", 4: "BMW"}
	var snacks = [...]string{"chips", "popcorn", "peanuts"}

	fmt.Println(zeroArray)
	fmt.Println(myArray)
	fmt.Println(cars)
	fmt.Println(snacks)

	// for length
	fmt.Println(len(snacks))

	// array is equal with same length and same elements and same order

	// slice
	someSlice := []int{5, 10, 15}
	fmt.Println(someSlice)

	// comparing the slice
	fruit1 := []string{"Apple", "Orange"}
	fruit2 := []string{"Apple", "Orange"}
	fruit1 = append(fruit1, "Mango")
	fruit3 := append(fruit1, fruit2...)
	fmt.Println(fruit3)
	fmt.Println(fruit1)
	fmt.Println(fruit3[1:2])
	fmt.Println(len(fruit1), len(fruit2))     // 2 , 2
	fmt.Println(slices.Equal(fruit1, fruit2)) // false

	values := []int{5, 6, 7, 8}
	newValues := make([]int, 4)
	copy(newValues, values)
	newValues[0] = 1337
	fmt.Println(newValues)
	fmt.Println(values)

}
