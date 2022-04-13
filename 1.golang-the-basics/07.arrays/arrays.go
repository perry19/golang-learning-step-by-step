package main

import "fmt"

func main() {
	// basic syntax: var arrayName [size] type

	var students [5]string
	fmt.Println(students[0]) // => since array is empty we have nothing printed

	// Let's assign a value to index 0
	// syntax array[index] = value

	students[0] = "Perry"
	fmt.Println(students[0]) // => Perry

	// To get the len of the array we use the built in len function
	fmt.Println("The size of the student array is:", len(students))

	// unlike in the previous steps we can directly initialize an array at declaration as follows
	team := []string{"Bob", "Martin", "Perry", "Poliatova", "Jinzhu", "Leonardo"}
	fmt.Println("Here is our team array:", team)

	// Multi-dimensional arrays
	// Below you will see how to define a 2-dimensional array

	var matrix [3][2]int

	for row := 0; row < 3; row++ {
		for col := 0; col < 2; col++ {
			matrix[row][col] = row + col
		}
	}
	fmt.Println("Matrix", matrix)

}
