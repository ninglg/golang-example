package main

import (
	"fmt"
	"os"
)

func osFunc() {
	fmt.Println("osFunc : ", os.Getenv("GOPATH"))
	// 变量os.Args是一个字符串slice
	// os.Args[0]是命令本身的名字
	// 实际参数范围是os.Args[1:len(os.Args)]
	fmt.Println("argsFunc os.Args : %+v", os.Args)
}

func main() {
	// os.Args
	osFunc()
}
