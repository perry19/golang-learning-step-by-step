package main

import (
	"fmt"
	"time"
)

func main() {
	println("==============Example one=================")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend 😁")
	default:
		fmt.Println("It's weekday 😒")
	}

	println("==============Example two=================")
	x := -1

	switch {
	case x % 2 == 0:
		fmt.Println("Even number")
	case x % 2 == 1:
		fmt.Println("Odd number")
	default:
		fmt.Println("Humm 🤔")		
	}
}
