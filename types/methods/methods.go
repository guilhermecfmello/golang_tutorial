package main

import (
	"fmt"
	"strings"
)

type people struct {
	firstName string
	lastName  string
}

func (p people) getFullName() string {
	return p.firstName + " " + p.lastName
}

func (p *people) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.firstName = parts[0]
	p.lastName = parts[1]
}

func main() {

	p1 := people{"Guilherme", "Mello"}

	fmt.Println(p1.getFullName())

	p1.setFullName("Admilson Jesus")
	fmt.Println(p1.getFullName())
}
