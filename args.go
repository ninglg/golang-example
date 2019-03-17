package main

import (
	"fmt"
)

func ts(args ...int) {
	fmt.Println("argsFunc : ", args)
}

func argsFunc() {
	ts(1, 2, 3)
	ts(4, 5)
	ts(6)
}
