package main

import (
	"fmt"
	"time"
)

func twoThreeFourMul(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	fmt.Println("Calculating C value...")
	time.Sleep(time.Second)
	c <- 4 * base
	fmt.Println("C Value calculated...")
}

func main() {
	c := make(chan int)
	fmt.Println("Before function call...")
	go twoThreeFourMul(9, c)
	fmt.Println("After function call!")

	fmt.Println("Getting value from channel...")
	a, b := <-c, <-c
	fmt.Println("Success getting values from channel!!")

	fmt.Printf("A: %d, B: %d\n", a, b)
	fmt.Printf("<-C: %d\n", <-c)
}
