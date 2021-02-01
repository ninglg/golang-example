package main

import "fmt"

// 递归函数
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func fibFunc() {
	fmt.Printf("fibonacci: ")

	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

	fmt.Println("")
}

func main() {
	fibFunc()
}
