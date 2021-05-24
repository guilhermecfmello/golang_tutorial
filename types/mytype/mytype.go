package main

import "fmt"

type grade float64

func (n grade) between(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func gradeToConcept(n grade) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 7.99) {
		return "C"
	}

	return "D"
}

func main() {
	var g1, g2, g3 grade
	g1 = 9.6
	g2 = 8.4
	g3 = 6.3
	fmt.Println("Grade ", g1, " concept: ", gradeToConcept(g1))
	fmt.Println("Grade ", g2, " concept: ", gradeToConcept(g2))
	fmt.Println("Grade ", g3, " concept: ", gradeToConcept(g3))
}
