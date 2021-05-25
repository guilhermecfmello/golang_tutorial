package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	// Block operation
	c <- 1
	fmt.Println("After reading")
}

func main() {
	c := make(chan int) // cannel without buffer

	go routine(c)
	// Block Operation
	fmt.Println(<-c)
	fmt.Println("Was read")
	// Deadlock
	fmt.Println(<-c)

	fmt.Println("Ending execution!")

}
