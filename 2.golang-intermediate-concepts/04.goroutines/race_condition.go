// Race condition

// A race condition occurs when the timing or order
// of events affects the correctness of a piece of code.

// consider the code below where we append marks to a slice
// The order of execution on the anonymous functions cannot be predicated
// and therefore the final order of values in the slice

// running this code with the -race flag (go run -race race_condition.go)
// Will show that there are some race conditions happening (In total 3 data races)

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Race Condition")
	var marks = []int{0}

	go func ()  {
		fmt.Println("Adding the 1st Mark")
		marks = append(marks, 1)
	}()

	go func ()  {
		fmt.Println("Adding the 2nd Mark")
		marks = append(marks, 2)
	}()

	go func ()  {
		fmt.Println("Adding the 3rd Mark")
		marks = append(marks, 3)
	}()

	time.Sleep(time.Millisecond * 50)

	fmt.Println("Final Marks : ", marks)
}