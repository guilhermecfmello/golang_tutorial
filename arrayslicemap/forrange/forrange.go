package main

import "fmt"

func main() {
	// Infering range of array
	numbers := [...]int{1, 2, 3, 4, 5}

	// range returns the index and the element respectively
	for i, number := range numbers {
		fmt.Printf("numbers[%d]: %d\n", i, number)
	}
	// Ignoring index with _ token
	for _, number := range numbers {
		fmt.Printf("numbers: %d\n", number)
	}
}
