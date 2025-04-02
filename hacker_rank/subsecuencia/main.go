package main

import "fmt"

func main() {
	fmt.Println(checkSubSecuence())
}

func checkSubSecuence() bool {
	// Program that checks if a slice is a subsequence of another slice
	arreglo := []int{5, 1, 22, 25, 6, -1, 8, 10}
	subsecuencia := []int{1, 6, -1, 10}
	isInArreglo := false

	for i := 0; i < len(subsecuencia); i++ {
		isInArreglo = false
		for j := 0; j < len(arreglo); j++ {
			if subsecuencia[i] == arreglo[j] {
				arreglo = arreglo[j:len(arreglo)]
				isInArreglo = true
				break
			}
		}
		if isInArreglo == false {
			return false
		}
	}
	return true
}
