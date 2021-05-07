package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// Switch without value looks for the first case that matches with a true condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good night")
	}
}
