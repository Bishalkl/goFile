package main

import "fmt"

// passing pointers to Function
func modify(p *int) {
	*p = *p + 10
}

// Structs & Pointers
type Person struct {
	Name string
	age  int
}

// function to modify the age
func (p *Person) birthday() {
	p.age++
}

func main() {
	x := 42
	p := &x // & Operator: address-of
	y := *p // * Operator: dereference
	*p = 24 // Modifying through pointer
	fmt.Println(p)
	fmt.Println(y)
	fmt.Println(*p)
	fmt.Println(x)

	// Pointer type declarations
	// var intPtr *int
	// var strPtr *string
	// var structPtr *struct{}

	z := 5
	modify(&z)
	fmt.Println(z)

	// pointer to struct
	p1 := &Person{"Bishal", 24}
	p1.birthday()
	fmt.Println(p1.age)
}
