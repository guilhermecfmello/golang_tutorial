package main

import "fmt"

type printable interface {
	toString() string
}

type people struct {
	firstName string
	lastName  string
}

type product struct {
	name  string
	price float64
}

func (p people) toString() string {
	return p.firstName + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var anything printable = people{"Jose", "Maria"}

	fmt.Println(anything.toString())
	print(anything)

	anything = product{"Calca Jeans", 79.90}
	fmt.Println(anything.toString())

}
