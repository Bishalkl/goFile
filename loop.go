package main

import (
	"fmt"
)

func main() {

	// creating for loop
	// for i:=0; i < 5; i++ {
	// 	fmt.Println(i)
	// }


	// creating while loop using forloop
	// i:=0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // inifinte loop 
	// for {
	// 	fmt.Println("This will run forever!")
	// }


	// Iterater Over a slice
	nums := []int {10, 20, 30, 40, 50}
	for index, value :=range nums {
		fmt.Println("Index:", index, "Value:", value)
	} 

	fmt.Println();

	// Iterate over a string (character-wise) use string()
	word := "Hello"
	for index, char := range word {
		fmt.Println("Index: ",index, "Character: ", string(char))
	}

	fmt.Println()

	// iterate over a map 
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	} 

	fmt.Println()


	// break statement
	for i:=0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println()


	// continue statement
	for i:=0; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()

	// goto statment
	i := 0
	START:
		fmt.Println(i)
		i++
		if i < 5 {
			goto START
		}
 
}

