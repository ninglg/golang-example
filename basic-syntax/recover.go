package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ~", r)
		// 打印出堆栈跟踪
		debug.PrintStack()
	}
}

func main() {
	// recover只能恢复同协程内的panic
	defer r()
	n := []int{7, 8, 9}

	// 产生一个运行时panic
	fmt.Println(n[3])
	fmt.Println("normally returned")
}
