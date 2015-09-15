package main

import (
	"fmt"
)

func main() {
	switch 5 + 4 {
	case 1:
		fmt.Println(1)
	case 2, 3, 4, 5, 6:
		fmt.Println(2, 3, 4, 5, 6)
	case 7, 8:
		fmt.Println(7, 8)
	case 9:
		fmt.Println(9)
		fallthrough
	case 0:
		fmt.Println(0)
	}
}
