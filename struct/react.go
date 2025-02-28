package main

import (
	"fmt"
)

// creating a type for rectangle
type rectangle struct {
	Width, Height float64
}

// Constructor function
func NewRectangle(width, height float64) rectangle {
	return rectangle{Width: width, Height: height}
}

// Method to calculate area
func (r rectangle) area() float64 {
	return r.Width * r.Height
}

// Method to calculate perimeter
func (r rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// main function
func main() {
	rect := rectangle{12.5, 34.5}
	fmt.Println("Area: ", rect.area())
	fmt.Println("Perimeter: ", rect.perimeter())

	rect1 := NewRectangle(12.4, 45.3)
	fmt.Println(rect1)
}
