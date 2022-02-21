package main

import "fmt"

// 如果变量没有明确初始化，它将隐式地初始化为这个类型的空值

func main() {
	a := "aaa"   //比较常用
	var b string //比较常用
	var c = "ccc"
	var d string = "ddd"

	fmt.Println(a, b, c, d)
}
