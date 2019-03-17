package main

import "fmt"

func forFunc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	for j, k := 0, 0; j+k < 100; j, k = j+5, k+10 {
		fmt.Printf("%d,%d ", j, k)
	}
	fmt.Println("")
}
