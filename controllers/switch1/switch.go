package main

import "fmt"

func gradeToConcept(n float64) string {
	var grade = int(n)

	switch grade {
	case 10:
		// To enter in other conditions
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade"
	}

}

func main() {
	var grade float64 = 3.4

	fmt.Println("Concept: " + gradeToConcept(grade))

}
