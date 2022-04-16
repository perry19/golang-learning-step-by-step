package main

import (
	"fmt"
)

func main() {
	// Starting with version 1.18, Go has added support for generics
	// also known as type parameters.

	// Generics allow our functions or data structures 
	// to take in several types that are defined in their generic form.

	// basically, They let us specify a function that can take in any kind of parameter.

	// Imagine writing a function max to return the maximum of two integers
	// You tested it worked well, but tommorrow you now want to implement thesame
	// thing, but this time for floats. What will you do ? 
	// I guess you will rewrite thesame by changing the types
	// instead of implementing thesame method over and over by changing only
	// The param types and return type, we can simply create a generic type to accept
	// params of any type
	fmt.Println("The max is", max(2.5, 7.3))
	fmt.Println("The max is", max(2, 7))
}

// Instead of using any, which is not comparable. we can use type set
// as shown below, to declare all the types our generic can accepts in an interface
type minTypes interface{
	~float64 | int
}

// In our function T is a generic type which can accept int and float 64
func max[T minTypes](a T, b T) T {
	if a > b {
		return a
	}
	return b
}