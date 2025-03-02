package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	time.Sleep(1000 * time.Millisecond) //Simulating some work
	fmt.Println("SayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hello Bishal")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Learning goroutines")
	go sayHello()
	go sayHi()

	// Wait for a moment to allow the goroutine to finish
	time.Sleep(3000 * time.Millisecond)
}
