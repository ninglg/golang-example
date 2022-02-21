package main

import "fmt"

// for是Go里面的唯一循环语句
// for循环的三个组成部分两边不用小括号，但大括号是必需的
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	for j, k := 0, 0; j+k < 100; j, k = j+5, k+10 {
		fmt.Printf("%d,%d ", j, k)
	}
	fmt.Println("")

	//for j < 10 {
	//	fmt.Println("这是while条件循环")
	//}

	//for {
	//	fmt.Println("这是死循环")
	//}
}
