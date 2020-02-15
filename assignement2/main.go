package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLenght float64
}

func main() {
	t := triangle{base: 2, height: 10}
	s := square{sideLenght: 5}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("The area is: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}
