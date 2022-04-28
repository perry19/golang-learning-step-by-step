package main

import "fmt"

func main() {
	// Slices are resizable arrays, and provide great flexibility
	// Therefore we can say a slice is an array without a fix size
	// This is how you declare a slice, just like an array: var aslice []string
	// or using the make function
	// When a slice is created it equally creates an underlying array...
	// In such cases the length and capacity of the slice are thesame
	
	// basic slice

	var mySlice = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice) // => [1 2 3 4 5]

	// creating a slice from another slice
	// slices support a “slice” operator with the syntax slice[low:high]
	var anotherSlice = mySlice[1:3]
	fmt.Println(anotherSlice) // => [2 3]

	// creating slices using the make function
	// Here we make a slice of strings of length 5 (initially zero-valued).
	var yourSlice = make([]string, 5)
	fmt.Println(yourSlice) // => [    ]

	// We can insert elements in the slice just like in arrays
	yourSlice[0] = "Martin"
	yourSlice[1] = "Job"

	fmt.Println(yourSlice) // => [Martin Job   ]

	// Use len() as with arrays to get the length of the slice
	// Use cap() to get the capacity of the slice: 
	// basically the capacity is the length of it underlying array 
	fmt.Println("Slice Length:", len(yourSlice))

	// slices support several more that make them richer than arrays. One is the builtin append
	fmt.Println("Updated slice is:", append(yourSlice, "Steve")) // => Updated slice is: [Martin Job    Steve]

	// Slices can also be copy’d.
	// Here we create an empty slice theirSlice of the same length as yourSlice
	// and copy into theirSlice from yourSlice.

	var theirSlice = make([]string, 5)
	copy(theirSlice, yourSlice)
	fmt.Println("The content of theirSlice is:", theirSlice) // => The content of theirSLice is: [Martin Job   ]

}
