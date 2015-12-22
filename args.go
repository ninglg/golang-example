package main

import (
	"fmt"
)

func test(arg ...int) {
	fmt.Println(arg)
}

func main() {
	test(1, 2, 3)
	test(4, 5)
	test(6)
}
