package main

import (
	"fmt"
)

func main() {
	var age int = 18
	if age < 18 && age >= 0 {
		fmt.Println("You ain't an adult yet")
	} else if age >= 18 {
		fmt.Println("You are an adult")
	} else{
		fmt.Println("We are confused, your age is surely negative :)")
	}
}