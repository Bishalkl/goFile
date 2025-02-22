package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd form painc:", r)
		}
	}()

	fmt.Println("start")
	panic("Something went wrong!")
	fmt.Println("End")
}
