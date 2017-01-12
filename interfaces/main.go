package main

import (
	"fmt"
	"math"
)

// Square is a shape
type Square struct {
	side float64
}

// Circle is a shape
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Shape is a shape
type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println("Area: ", s.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
