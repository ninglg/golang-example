package main

import (
	"fmt"
)

func test(x int) (y int, z int) {
	y = x
	z = x * 2
	return
}

func main() {
	fmt.Println(test(6))
}
