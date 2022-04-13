package main

import "fmt"

func main() {
	// range iterates over elements in a variety of data structures.

	// Let's create a slice containing some numbers and iterate through it to get the sum

	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum) // => 15

	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.

	for index, value := range numbers {
		if value == 4 {
			fmt.Println("The index of 4 is:", index) // => The index of 4 is: 3
		}
	}

	// range on map iterates over key/value pairs.
	var countries = make(map[string]string)
	countries["JP"] = "Japan"
	countries["CMR"] = "Cameroon"
	countries["RU"] = "Russia"

	for k, v := range countries {
		fmt.Println(k, v)
	}
}
