package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j, k := 0, 0; j+k < 100; j, k = j+5, k+10 {
		fmt.Println(j, k)
	}
}
