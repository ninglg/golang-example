package main

import "fmt"

func switchFunc() {
	var ss int
	ss = 10
	switch ss {
	case 0:
		fmt.Println("case 0")
	case 10:
		fmt.Println("case 10")
		fallthrough
	case 20:
		fmt.Println("case 20 fallthrough")
	default:
		fmt.Println("case default")
	}
}
