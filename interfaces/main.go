package main

import "fmt"

func main() {
	triang := triangle{height: 2, base: 3}
	sq := square{sideLength: 5}

	printArea(triang)
	printArea(sq)
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
