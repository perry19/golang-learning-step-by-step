package main

import (
	"fmt"
)

// A channel is simply a way for goroutines to communicate with each other.
// They can be thought of as pipes which are used by goroutines to communicate.
// Each channel variable can hold data only of a particular type.
// Channels are declared using the chan keyword
// the default value for a channel is nil
func main() {
	// channel declarartion
	var a chan int

	// output -> <nil>
	fmt.Println(a)

	// define a channel using the make keyword
	a = make(chan int)

	// output -> 0xc000020120 (Address might be different on your machine)
	fmt.Println(a)

	// There are 2 main operations that can be done on channels: Send and Receive
	// Send operation is used to send data to the channel and has the syntax: ch <- val
	// ch is the channel name and val is the value being sent to the channel

	// The receive operation is used to read data from a channel. Syntax: ch := <- val
	
}