package main

import "fmt"

func main() {
	// A pointer references a location in memory
	// obtaining the value stored at that location is known as dereferencing the pointer

	// We initialize a variable p with the value 30
	p := 30

	// The &p stands for address of p, meaning we are passing the adress of p to the variable q
	q := &p

	fmt.Println("The address of p is:", q) // => The address of p is: 0xc000006028
	fmt.Println("The value of p is:", *q)  // => The value of p is: 30

	// notice the type of q, an integer pointer
	fmt.Printf("q is of type %T\n", q) // => q is of type *int

	// Since q points to p, changing it's value equally changes the value of p
	// instead of making a copy
	*q = 20
	fmt.Println("P is now:", p) // => P is now: 20

	// why do we need pointers
	// It's more efficient to store and manipulate the variable in one location
	// Rather than copying the variables anytime you need to make changes
	// However one needs to be careful when working with pointers...
	// Can read more on this online
}
