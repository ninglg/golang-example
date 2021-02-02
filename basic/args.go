package main

import (
	"fmt"
)

func ts(args ...int) {
	fmt.Println("argsFunc : ", args)
}

func argsFunc() {
	ts(1, 2, 3)
	a := []int{4, 5, 6}
	ts(a...)
	ts(7, 8)
}
