package main

import "fmt"

// 声明全局变量
var (
	name string
	age  int
	sex  byte // 全局变量允许声明但不使用
)

func main() {
	name = "xiaowang"
	age = 16

	fmt.Println(name, age)
}
