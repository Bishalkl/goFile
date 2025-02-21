package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int64
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	// to print strign value 
	var myString string = "Hello \nWorld"
	var myString1 string = "Hello world"
	fmt.Println(myString)
	fmt.Println(myString1)

	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// fancy way intialize the value is 
	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)
	
	// const value 
	const myConst string = "const value"
	fmt.Println(myConst)
}

