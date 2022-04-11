package main

import "fmt"

// In golang there's only a single loop construct: The for loop
// However, it can be implemented in a few ways

func main() {

	// classical for loop
	fmt.Println("classical for loop")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// While loop-like variation
	fmt.Println("While loop-like variation")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Others
	fmt.Println("We will look other techniques When we look at range statements, channels, and other data structures.")
}
