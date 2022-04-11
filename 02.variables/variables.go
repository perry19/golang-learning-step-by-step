package main

import "fmt"

func main() {

	// long declaration
	var name string = "Bob"
	fmt.Println("Hello Mr", name)

	// short declaration
	surname := "Alice"
	fmt.Println("Hello Mrs", surname)

	// declare many variables at once
	// Here we are using a printing variation similar to what you see in languages like C and C++

	var a, b int = 2, 4
	fmt.Printf("The value of a is: %d\n", a)
	fmt.Printf("The value of b is: %d\n", b)
}