package main

import (
	"fmt"
	"os"
)

func modify() (x int) {
	x = 10 //not using the : becasue variable value gonna be change
	defer func() {x += 5}()
	return
} 

func main() {
	file, err := os.Create("test.txt")
	if err!= nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Ensures file closes after function execution
	file.WriteString("Hello, go")

	fmt.Println()

	fmt.Println(modify()) //Output: 15
}