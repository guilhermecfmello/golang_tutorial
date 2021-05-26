package main

import (
	"fmt"
	"time"

	htmlProcessor "github.com/guilhermecfmello/html"
)

func theFaster(url1, url2, url3 string) string {
	c1 := htmlProcessor.Title(url1)
	c2 := htmlProcessor.Title(url2)
	c3 := htmlProcessor.Title(url3)

	// Control structure for concurrency
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Millisecond * 1000):
		return "Everyone lost"
		// default:
		// 	return "Still no response"
	}
}

func main() {
	winner := theFaster(
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.estadao.com.br",
	)

	fmt.Println(winner)
}
