package main

import "fmt"

func sliceFunc() {
	se := make([]int, 10, 20)
	fmt.Println("len(se) = ", len(se))
	fmt.Println("cap(se) = ", cap(se))
	se = []int{5, 6, 7, 8}
	fmt.Println("se = ", se)
	se = append(se, 9)
	fmt.Println("append se = ", se)
	fmt.Println(se[1:3])
}
