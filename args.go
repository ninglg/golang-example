package main

import (
	"fmt"
)

func ts(arg ...int) {
	fmt.Println(arg)
}

func main() {
	ts(1, 2, 3)
	ts(4, 5)
	ts(6)
}
