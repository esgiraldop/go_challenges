package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateInput(num interface{}) (int, bool) {
	// Returns false if the input is invalid
	var message strings.Builder
	switch v := num.(type) {
	case string:
		val, err := strconv.Atoi(v)
		message.WriteString("It is a string")
		if err != nil {
			message.WriteString(", and cannot be converted to an integer.")
			fmt.Println(message.String())
			return val, false
		} else {
			message.WriteString(", but can be converted to an integer value.")
			if val < 1 || val > 11 {
				message.WriteString(" The length of the vector must be between 1 and 11")
				return val, true
			}
			fmt.Println(val)
			fmt.Println(message.String())
			return val, true
		}
	case int:
		message.WriteString("It is an integer.")
		fmt.Println(message.String())
		return v, true
	default:
		message.WriteString("It is neither a string nor an integer.")
		fmt.Println(message.String())
		return 0, false
	}
}

func newArray(inputLength interface{}) []int {
	// Function that returns an array of length "length" with random integer numbers between 1 and 10. 1 < length <= 11
	length, isValid := validateInput(inputLength)
	fmt.Println("answer: ", length)
	fmt.Println("isValid: ", isValid)
	// if isValid {
	// 	var intArray []int

	// 	for i := 0; i < length; i++ {

	// 	}
	// }
	dummyArray := []int{1, 2}
	return dummyArray
}
