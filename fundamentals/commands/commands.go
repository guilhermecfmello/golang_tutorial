package main

import "fmt"

func main() {
	fmt.Printf("Another %s program!", "GO")
}

// go doc cmd/vet: identify unreachable code
// go vet [FILENAME]: verify warnings on go code
// go build [FILANAME], then ./FILENAME: Compile and build
// go run commands.go
//
