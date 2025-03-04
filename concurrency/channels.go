package main

import "fmt"

func main() {
	// creating channels
	messageChan := make(chan string)

	// sending data
	messageChan <- "ping" // blocking

	msg := <-messageChan

	fmt.Println(msg)
}
