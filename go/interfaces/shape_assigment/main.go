package main

import "fmt"

type square struct {
	sizeLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sizeLength: 4}
	t := triangle{height: 5, base: 2}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sizeLength * s.sizeLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}
