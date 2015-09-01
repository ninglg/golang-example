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

	const d = iota

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
