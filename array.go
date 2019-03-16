package main

import "fmt"

func arrayFunc() {
	arr1 := [3]int{22, 33, 44}
	arr2 := [...]string{"hello", "world"}
	fmt.Println(arr1)
	fmt.Println(arr2)
}
