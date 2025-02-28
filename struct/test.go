package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method with a value receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with a pointer receiver to modify the struct
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	// Using field names
	p1 := Person{Name: "Bishal", Age: 23}

	// Without field names (order matters)
	p2 := Person{"Alex", 30}

	// Using var (default zero values)
	var p3 Person

	fmt.Println(p1) // Output: {Bishal 22}
	fmt.Println(p2) // Output: {Alex 30}
	fmt.Println(p3) // Output: {0} (default values)

	p := Person{Name: "Bishal koirala", Age: 22}
	p.Greet()

	p4 := Person{Name: "Komal koirala", Age: 18}
	p4.HaveBirthday()
	fmt.Println(p4.Age)

}
