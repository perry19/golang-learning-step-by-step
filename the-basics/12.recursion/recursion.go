package main

import "fmt"

func main() {
	fmt.Println(fact(5))
	fmt.Println(fib(7))
}

// A recursive function is a function that calls it self
// This fact function calls itself until it reaches the base case of fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//The Fibonacci numbers, commonly denoted F(n) form a sequence
//called the Fibonacci sequence, such that each number is the sum of the two preceding ones
//starting from 0 and 1. That is,

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).
func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}