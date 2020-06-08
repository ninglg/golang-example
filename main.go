// go run *.go
package main

// 分组顺序：系统包、项目内部包、外部包
import (
	"fmt"
)

func main() {
	// main之前有2个init函数被自动调用

	// 开始执行main内部逻辑
	fmt.Println("...main函数开始...")

	// os包及os.Args
	osFunc()

	// i++是语句，不是表达式，所以 j = i++ 是不合法的
	// 仅支持后缀形式，所以 --i 也是不合法的

	// for 是 Go 里面唯一的循环语句

	// 可变参数
	argsFunc()
	// flag
	flagFunc()

	// invalid
	invalidFunc()
	// time
	timeFunc()
	// import
	importFunc()
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
	// 随机数
	randFunc()
	// 字符串
	stringFunc()
	// 结构体 + 方法
	circleFunc()
	// 数组
	arrayFunc()
	// 结构体
	structFunc()
	// 切片
	sliceFunc()
	// map + range + ok + delete
	mapFunc()
	// 递归函数
	fibFunc()
	fibFunc2()
	// goto
	gotoFunc()
	// 类型转换
	transFunc()
	// 排序
	sortFunc()
	// Type
	typeFunc()
	// 接口
	infFunc()
	// 错误处理
	errorFunc()
	// Recover
	recoverFunc()

	// 不要以共享内存的方式来通信，而要以通信来共享内存。
	// CSP：通信顺序进程
	// 并发
	goFunc()
	// 通道
	channelFunc()
	// json
	jsonFunc()

	// http
	//httpFunc()
	// 并发http
	concurrenceHTTP()
	// file server
	//fileServerFunc()
	// http server
	//httpServerFunc()
	pipelineFunc()
	// 文件操作
	fileFunc()
	// Base64
	base64Func()

	fmt.Println("...main函数结束...")
}
