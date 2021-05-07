package main

import "fmt"

// There is not ternary operator in Go
// This is a solution
func getResult(grade float64) string {
	if grade >= 6 {
		return "Approved"
	}

	return "Disapproved"
}

func main() {
	var grade float64 = 6.2

	fmt.Println("Grade: ", getResult(grade))
}
