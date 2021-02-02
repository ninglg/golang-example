package main

import (
	"fmt"
)

var done bool

func test() {
	fmt.Println("test")
	done = true
}

func main() {
	go test()

	for !done {

	}
}