package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 1)
}

func send(ch chan string) {
	name := "Jack"
	ch <- name
}

func receive(ch chan string) {
	val := <-ch
	fmt.Printf("value received = %s", val)
}