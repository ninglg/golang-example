package main

import "fmt"

// 如果变量没有明确初始化，它将隐式地初始化为这个类型的空值

// 可用分组方式定义变量
const (
	i      = 100
	pi     = 3.14
	prefix = "Go_"
)

var (
	a int
	c string
)

func main() {
	a := "aaa"   //比较常用
	var b string //比较常用
	var c = "ccc"
	var d string = "ddd"

	fmt.Println(a, b, c, d)

	// 对于接口和引用类型（slice、map、指针、channel、func）的初始值是nil
	// Go里面不存在未初始化变量

	// 函数返回局部变量的地址是非常安全的

	// rune类型是int32类型的同义词，byte类型是unit8类型的同义词
	// 无符号整数uintptr，其大小并不明确，但足以完整存放指针
	// Go具有两种大小的浮点数float32和float64
	// Go具备两种大小的复数complex64和complex128

	// 原生的字符串字面量，适用于HTML模板、Json字面量、命令行提示信息、以及需要多行文本表达的场景
	str := `
	hello, world.
	this is a test.
	`
	fmt.Println(str)
}
