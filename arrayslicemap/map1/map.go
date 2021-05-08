package main

import "fmt"

func main() {
	// var approved map[int]string
	// maps must be inicialized
	approved := make(map[int]string)
	approved[123456789] = "Gui"
	approved[129391293] = "Noir"
	approved[1] = "God"

	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Println(cpf, name)
	}

	fmt.Println(approved)
	delete(approved, 123456789)
	fmt.Println(approved)

}
