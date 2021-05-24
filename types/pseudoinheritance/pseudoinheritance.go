package main

import "fmt"

type car struct {
	name  string
	speed int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.speed = 220
	f.turboOn = true

	fmt.Printf("Ferrari %s is with turbo on? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}
