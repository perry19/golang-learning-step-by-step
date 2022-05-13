package main

import (
	"fmt"
	"time"
)

// A goroutine is a light-weight thread managed by the go runtime

// before talking of goroutines let's understand what is concurrency and parallelism

// In laymen term concurrency is the

func sayHi(name string) string {
	return "Hello " + name
}

func main() {
	go fmt.Println(sayHi("Mary"))

	fmt.Println(sayHi("John"))

	time.Sleep(time.Microsecond)

	fmt.Println("Done!")
}

