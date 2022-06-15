package main

import (
	"fmt"
	"time"
)

// A goroutine is a light-weight thread managed by the go runtime

// before talking of goroutines it will be good to understand what is concurrency and parallelism

// To invoke a go routine use the go keyword followed by the function e.g go f(s)

func sayHi(name string) string {
	return "Hello " + name
}

func main() {
	go fmt.Println(sayHi("Mary"))

	fmt.Println(sayHi("John"))

	time.Sleep(time.Microsecond)

	fmt.Println("Done!")
}
