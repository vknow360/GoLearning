package main

import (
	"fmt"
	"math"
)

// An interface in go defines a set of method signatures.
// Any type that implements those methods satisfies the interface.

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var g geometry

	g = rectangle{width: 3, height: 4}
	println("Rectangle Area:", g.area())
	println("Rectangle Perimeter:", g.perimeter())

	g = circle{radius: 5}
	println("Circle Area:", g.area())
	println("Circle Perimeter:", g.perimeter())

	myPrinter(1, "John", 45.9, true)
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
