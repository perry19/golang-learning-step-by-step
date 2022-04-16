package main

import "fmt"

func main() {
	// Go supports methods defined on struct types.
	r := Rect{9, 5}
	fmt.Println("Area:", r.calculateArea())

	fmt.Println("Perimeter", r.calculatePerimeter())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls
	// or to allow the method to mutate the receiving struct.

	rp := &r

	fmt.Println("Area:", rp.calculateArea())
	fmt.Println("Perimeter:", rp.calculatePerimeter())
}

// Lets create a method to calculate the area of a rectangle
// A rectangle will be defined by a structure Rect

type Rect struct {
	l uint
	w uint
}

// Here is the syntax for a method.
// func (receiver type) method() returnType {

// }
// This calculateArea method has a receiver type of Rect.
func (rect Rect) calculateArea() uint {
	area := rect.l * rect.w
	return area
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a pointer receiver.

func (rect *Rect) calculatePerimeter() uint {
	return 2 * (rect.l + rect.w)
}
