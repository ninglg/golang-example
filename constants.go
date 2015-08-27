package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = 3
		c = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
