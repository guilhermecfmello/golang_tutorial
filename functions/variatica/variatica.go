package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}

func main() {
	n1, n2, n3 := 12.0, 8.0, 5.0
	fmt.Println(average(n1, n2, n3))
}
