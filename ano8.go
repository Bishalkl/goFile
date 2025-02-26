package main

import "fmt"

func idGenerator() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

func main() {
	gen := idGenerator()

	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	
}

