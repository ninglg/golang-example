package main

import (
	"fmt"
)

// init函数在main之前执行
func init() {
	fmt.Println("...第1个init函数...")
}

// 可以定义多个init函数
func init() {
	fmt.Println("...第2个init函数...")
}

func initFunc() {
	//init自动执行
}