package main

import (
	"fmt"
	"time"
)

func talk(people, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second / 10)
		fmt.Printf("%s: %s (iteration %d)\n", people, text, i+1)
	}
}

func main() {
	go talk("DJ GBR", "Eh o brabo de novo", 300)
	talk("DJ Cleiton Rasta", "Cabeca de gelo", 200)

	// time.Sleep(time.Second * 5)
	fmt.Println("End")
}
