package main

import (
	"fmt"
	"log"
)

func main() {
	// get user input into slices
	var count int = 10
	values := make([]int, 10)
	for i := 0; i < count; i++ {
		fmt.Printf("Enter the value for value %d: ", i + 1)
		value, err := fmt.Scan(&values[i])
		if err != nil{
			log.Fatal(err)
		}
		values = append(values, value)
	}
	fmt.Println("Sorted slice")
	bubbleSort(values[:10])
}

func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j + 1] {
				Swap(arr, j)
			}
		}
	}
	fmt.Println("\nAfter Bubble Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}

func Swap(nums []int, i int) {
	tmp := nums[i]
	nums[i] = nums[i+1]
	nums[i+1] = tmp
}