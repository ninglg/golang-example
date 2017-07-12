package main

import (
	"fmt"
)

func ff(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func main() {
	fmt.Println(ff(0))
}