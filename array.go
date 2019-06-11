package main

import "fmt"

func arrayFunc() {
	fmt.Println("========== arrayFunc ==========")
	arr1 := [3]int{22, 33, 44}
	arr2 := [...]string{"hello", "world"}
	arr3 := [2][3]int{{4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
