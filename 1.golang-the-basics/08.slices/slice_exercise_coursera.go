/*
	Write a program which prompts the user to enter integers and stores
	the integers in a sorted slice.The program should be written as a loop.
	Before entering the loop, the program should create an empty integer slice
	of size (length) 3. During each pass through the loop, the program prompts
	the user to enter an integer to be added to the slice. The program adds the
	integer to the slice, sorts the slice, and prints the contents of the slice
	in sorted order. The slice must grow in size to accommodate any number
	of integers which the user decides to enter. The program should only quit
	(exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var nums = make([]string, 3)
	nums = nil
	var ch string

	fmt.Println("Instructions: Enter a sequence of integers, press the enter key after entering each integer")
	fmt.Println("When you are done Entering the integers type X to exit. Let's begin...")
	fmt.Println("")
	fmt.Println("Enter a sequence of integer:")

	i := 0

	for {
		fmt.Scanf("%s", &ch)
		nums = append(nums, ch)
		i++
		if ch == "X" || ch == "x" {
			break
		}
	}
	ints := make([]int, len(nums))

	for i, s := range nums {
		ints[i], _ = strconv.Atoi(s)
	}
	sort.Ints(ints[:])
	fmt.Println(ints[1:])
}
