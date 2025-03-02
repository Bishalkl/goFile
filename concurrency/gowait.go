package main

import (
	"fmt"
	"sync"
	"time"
)

// Task one
func taskOne(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Task one is done!")
}

// Task two
func taskTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task two is done!")
}

// Task three
func taskThree(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Task three is done!")
}

// Task four
func taskFour(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("Task four is done!")
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("Start the task: ")

	wg.Add(4)
	go taskOne(&wg) //add the reference and pass it
	go taskTwo(&wg)
	go taskThree(&wg)
	go taskFour(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed!")
}
