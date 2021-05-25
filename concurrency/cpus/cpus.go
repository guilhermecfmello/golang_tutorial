package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU qtd of this machine: ", runtime.NumCPU())
}
