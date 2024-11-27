package main

import "fmt"

func main() {
	length, isValid := validateInput(10)
	fmt.Println("isValid: ", isValid)
	if isValid {
		intArray := newArray(length)
		checkEven(intArray)
	}
}
