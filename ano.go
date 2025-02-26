// Write an anonymous function that prints "Hello, Go!" and invoke it immediately.
package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello, Go!")
	}()

}
