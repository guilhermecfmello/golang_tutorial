package main

import (
	"fmt"

	htmlProcessor "github.com/guilhermecfmello/html"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

// Multiplex: Mixing data from different channels
func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(input1, c)
	go forward(input2, c)
	return c
}

func main() {
	c := join(
		htmlProcessor.Title("https://www.cod3r.com.br", "https://www.globo.com"),
		htmlProcessor.Title("https://www.youtube.com.br", "https://estadao.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
