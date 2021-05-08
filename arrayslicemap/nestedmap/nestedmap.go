package main

import "fmt"

func main() {

	funcsByLetter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  123123.123,
			"Guilherme Mello": 2300.99,
		},
		"J": {
			"Jose Camargo": 987.90,
		},
		"Z": {
			"Zuraide Pereira": 1230.00,
			"Ze do mato":      9999.99,
		},
	}

	delete(funcsByLetter, "G")

	for letter, funcs := range funcsByLetter {
		fmt.Printf("Letter: %s\n", letter)
		for name, salary := range funcs {
			fmt.Printf("\tName: %s\n\tSalary: %v\n", name, salary)
		}
	}
}
