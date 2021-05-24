package main

import "fmt"

type sportive interface {
	turnOnTurbo()
}

type ferrari struct {
	model   string
	turboOn bool
	speed   int
}

func (f *ferrari) turnOnTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()
	fmt.Println(car1)

	var car2 sportive = &ferrari{"F40", false, 0}
	fmt.Println(car2)

}
