// Write a function that returns a function to greet users in different languages, such as "Hello" (English), "Hola" (Spanish), and "Bonjour" (French).
package main

import "fmt"

func main() {
	greet := func() func(string) string {
		return func(language string) string {
			if language == "English" {
				return "Hello"
			} else if language == "Spanish" {
				return "Hola"
			} else if language == "French" {
				return "Bonjour"
			} else {
				return "Sorry, don't know this language!"
			}
		}
	}() // invoke it self

	// result
	fmt.Println(greet("English"))
	fmt.Println(greet("Spanish"))
	fmt.Println(greet("French"))
	fmt.Println(greet("Nepali"))
}
