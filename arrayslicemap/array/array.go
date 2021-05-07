package main

import "fmt"

func main() {

	// homogenea (same data type) and static
	// definition:
	var grade [3]float64

	fmt.Println(grade)

	grade[0], grade[1], grade[2] = 5.6, 5, 1.4
	fmt.Println(grade)

	total := 0.0
	for i := 0; i < len(grade); i++ {
		total += grade[i]
	}

	average := total / float64(len(grade))

	fmt.Printf("Average: %.2f", average)
}
