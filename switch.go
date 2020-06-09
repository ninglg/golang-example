package main

import "fmt"

func switchFunc() {
	var ss int
	ss = 10
	switch ss {
	case 0:
		fmt.Println("case 0")
	case 10, 20, 30, 40:
		fmt.Println("case 10")
		fallthrough
	case 50:
		fmt.Println("case 50 fallthrough")
	default:
		fmt.Println("case default")
	}

	//
	switch 5 + 4 {
	case 1:
		fmt.Println(1)
	case 2, 3, 4, 5, 6:
		fmt.Println(2, 3, 4, 5, 6)
	case 7, 8:
		fmt.Println(7, 8)
	case 9:
		fmt.Println("case : ", 9)
		fallthrough
	case 0:
		fmt.Println(0)
	}

	//
	switch {
	case 5 > 10:
		fmt.Println("case : 5 > 10")
	case 4 < 8:
		fmt.Println("case : 4 < 8")
	}
}

func main() {
	switchFunc()
}
