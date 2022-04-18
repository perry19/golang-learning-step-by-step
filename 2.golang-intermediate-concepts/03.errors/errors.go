package main

import (
	"errors"
	"fmt"
)

// In Go, errors are the last return value by convention
// and have type error, a built-in interface.
// That means the error can be checked immediately before proceeding to the next steps.

// Simple error example
// consider the below method. We will return an error in case our number is negative

func checkNumber(num int) (int, error) {
	if num < 0 {
		return num, errors.New("The provided number is negative")
	}
	return num, nil
}

// We can equally create custom errors by implementing the error interface
// The error interface contains a single method called Error that returns a string

type customError struct {
	errorMsg string
	arg      int
}

func (e *customError) Error() string {
	return fmt.Sprintf("%s - %d", e.errorMsg, e.arg)
}

func checkNumberV2(arg int) (int, error) {
	if arg < 0 {
		return -1, &customError{"Error: The error occured due to", arg}
	}
	return arg, nil
}

// To check for an error we simply get the second value of the function and then
// check the value with the nil. Since the zero value of an error is nil.
// So, we check if an error is a nil.
// If it is then no error has occurred and all other cases the error has occurred.
func main() {

	// Testing simple error example
	result, err := checkNumber(-4)

	if err != nil {
		fmt.Println(result, err) // => -4 The provided number is negative
	} else {
		fmt.Println(result)
	}

	// Testing custom error example
	result1, err1 := checkNumberV2(-3)

	if err1 != nil {
		fmt.Println(result1, err1) // => -1 Error: The error occured due to - -3
	} else {
		fmt.Println(result1)
	}

}
