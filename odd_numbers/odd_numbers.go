package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func validateInput(num interface{}) (int, bool) {
	// Returns false if the input is invalid
	var message strings.Builder
	switch v := num.(type) {
	case string:
		val, err := strconv.Atoi(v)
		message.WriteString(fmt.Sprintf("%d is a string", num))
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
		message.WriteString(fmt.Sprintf("%d is an integer", num))
		fmt.Println(message.String())
		return v, true
	default:
		message.WriteString(fmt.Sprintf("%d is neither a string nor an integer.", num))
		fmt.Println(message.String())
		return 0, false
	}
}

func newArray(inputLength int) []int {
	// Function that returns an array of length "length" with random integer numbers between 1 and 10. 1 < length <= 11
	var intArray []int
	for i := 0; i < inputLength; i++ {
		RandomIntegerwithinRange := rand.Intn(100-1) + 1
		intArray = append(intArray, RandomIntegerwithinRange)
	}
	fmt.Println("The array is: ", intArray)
	return intArray
}

func checkEven(intArray []int) {
	for _, num := range intArray {
		if num%2 == 0 {
			fmt.Println(fmt.Sprintf("%d is even", num))
		} else {
			fmt.Println(fmt.Sprintf("%d is odd", num))
		}
	}
}
