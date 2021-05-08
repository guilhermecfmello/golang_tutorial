package main

import "fmt"

func main() {

	funcsAndSalaries := map[string]float64{
		"Joao Souza": 1231.12,
		"Gui Mello":  1800.12,
		"Ane Silva":  2312.13,
	}

	funcsAndSalaries["Gui Mello"] = 999.0
	fmt.Println(funcsAndSalaries)

	// There is no problem trying to exclude not existent values
	delete(funcsAndSalaries, "aaa")

	for name, salary := range funcsAndSalaries {
		fmt.Printf("Name: %s || Salary: %.2f\n", name, salary)
	}
}
