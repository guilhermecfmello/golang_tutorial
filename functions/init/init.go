package main

import "fmt"

// init is executed before main as convenction
func init() {
	fmt.Println("starting")
}

func main() {
	fmt.Println("Main...")
}
