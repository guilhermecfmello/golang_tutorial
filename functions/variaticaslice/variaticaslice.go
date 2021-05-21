package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("approved list: ")
	for i, approved := range approveds {
		fmt.Printf("[%d]: %s\n", i, approved)
	}
}

func main() {
	approveds := []string{"Maria", "epaminondas", "up altas aventuras", "Gui"}
	// Separates elements of slice in various elements of parameters
	printApproveds(approveds...)
}
