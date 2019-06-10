package main

import (
	"fmt"
	"os"
)

func ts(args ...int) {
	fmt.Println("argsFunc : ", args)
}

func argsFunc() {
	ts(1, 2, 3)
	a := []int{4, 5, 6}
	ts(a...)
	ts(7, 8)
	fmt.Println("argsFunc os.Args : ", os.Args)
}
