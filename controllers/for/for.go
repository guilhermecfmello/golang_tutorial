package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// for is the only repetition control structure
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("I: %d\n", i)
	}

	fmt.Println("Mixing...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// infinity loop
	for {
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}
}
