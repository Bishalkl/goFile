package main

import (
	"fmt"
)

func main() {
	// Loop to defer printing numbers
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
