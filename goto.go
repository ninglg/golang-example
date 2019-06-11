package main

import "fmt"

func gotoFunc() {
	fmt.Println("========== gotoFunc ==========")
	fmt.Printf("gotoFunc : ")
	i := 0
Loop:
	if i < 10 {
		fmt.Printf("%d ", i)
		i++
		goto Loop
	}
	fmt.Println("")
}
