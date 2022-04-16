package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.

// Let's create a basic interface for geometric shapes.

type geometry interface {
	area() float64
	perimeter() float64
}

// For our example we’ll implement this interface on rect and circle types.

type rect struct {
	length, width float64
}

type circle struct {
	radius float64
}

// To implement an interface, we just need to implement all the methods in the interface.
// Here we implement geometry on rects.

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.length + r.width
}

// And here we implement geometry on circle

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	r := rect{5, 3}
	c := circle{5}
	fmt.Println("The Area of the Rectangle is:", r.area())
	fmt.Println("The Perimeter of the Rectangle is:", r.perimeter())
	fmt.Println("The Area of the Circle is:", c.area())
	fmt.Println("The Perimeter of the Circle is:", c.perimeter())
}
