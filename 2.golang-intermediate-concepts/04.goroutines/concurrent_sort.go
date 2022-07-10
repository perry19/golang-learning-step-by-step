package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

    fmt.Println("Enter a series of space seperated integer: ")

    scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	var n []int  

	for _, token := range strings.Split(userInput, " ") {
		number, _ := strconv.Atoi(token)
		n = append(n, number)
	}

	fmt.Println("Initial Array: ", n)

	fmt.Println("splitted Array: ", splitArray(n, 4))
	splittedArray := splitArray(n, 4)

	var wg sync.WaitGroup


	wg.Add(4)
	go func() {
		defer wg.Done()
		sortArray(splittedArray[0])
	}()
	go func() {
		defer wg.Done()
		sortArray(splittedArray[1])
	}()
	go func() {
		defer wg.Done()
		sortArray(splittedArray[2])
	}()
	go func() {
		defer wg.Done()
		sortArray(splittedArray[3])
	}()

	wg.Wait()

	fmt.Println("sorted sub arrays: ", splittedArray)

	fmt.Println("Final sorted array: ", sortArray(mergeArray(splittedArray)))
}

func sortArray(n []int) []int {
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}

	return n
}

func splitArray(array []int, numberOfChunks int) [][]int {
	if len(array) == 0 {
		return nil
	}
	if numberOfChunks <= 0 {
		return nil
	}

	if numberOfChunks == 1 {
		return [][]int{array}
	}

	result := make([][]int, numberOfChunks)

	// we have more splits than elements in the input array.
	if numberOfChunks > len(array) {
		for i := 0; i < len(array); i++ {
			result[i] = []int{array[i]}
		}
		return result
	}

	for i := 0; i < numberOfChunks; i++ {

		min := (i * len(array) / numberOfChunks)
		max := ((i + 1) * len(array)) / numberOfChunks

		result[i] = array[min:max]

	}

	return result
}

func mergeArray(arr [][]int) []int {
	var n []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			n = append(n, arr[i][j])
		}
	}
	return n
}
