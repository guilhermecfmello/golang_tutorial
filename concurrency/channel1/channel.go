package main

import "fmt"

// Channel is a stack shared with other processes

func main() {
	// Creating channel
	ch := make(chan int, 1)
	// Sending data to the channel (writing)
	ch <- 1
	// Receiving data from channel
	<-ch

	ch <- 2
	fmt.Println(<-ch)

}
