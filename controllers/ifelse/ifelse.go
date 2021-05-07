package main

import "fmt"

func printResult(grade float64) {
	// There is note parentesis in if
	if grade >= 7 {
		fmt.Println("Approved with grade: ", grade)
	} else {
		fmt.Println("Disapproved with grade: ", grade)
	}
}

func printResultAsConcept(grade float64) string {
	// There is note parentesis in if
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade <= 9 {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	var grade float64 = 5.7

	printResult(grade)
	fmt.Println("Result as concept: " + printResultAsConcept(grade))
}
