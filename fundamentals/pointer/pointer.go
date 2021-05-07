package main

import "fmt"

func main() {
	i := 1

	// There isn't arithmetic operations with pointers
	var p *int = nil

	p = &i

	fmt.Println("I as P: ", *p)
	fmt.Println("I address: ", p)

	//  Increment as pointer
	*p++
	fmt.Println("I as P: ", *p)
	fmt.Println("I address: ", p)
}
