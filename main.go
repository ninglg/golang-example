// go run *.go
package main

import (
	"fmt"
)

func main() {
	// 这里是注释
	fmt.Println("...main函数开始...")

	// 全局变量
	nameFunc()
	// 短声明只能用于函数体内
	colorFunc()
	// 函数返回值及空白标识符
	dummyFunc()
	// 常量
	infoFunc()
	// iota
	iotaFunc()
	// switch
	switchFunc()
	// for
	forFunc()
	// 函数作为值
	squareFunc()
	// 结构体 + 方法
	circleFunc()
	// 数组
	arrayFunc()
	// 切片
	sliceFunc()
	// map + range + ok + delete
	mapFunc()
	// 递归函数
	fibFunc()
	// 类型转换
	transFunc()
	// 接口
	infFunc()
	// 错误处理
	errorFunc()
	// 并发
	goFunc()
	// 通道
	channelFunc()

	fmt.Println("...main函数结束...")
}
