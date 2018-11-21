package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = 3
		c = iota
		d = 1 << iota
	)

	const e = iota

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
