package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("PATH : ", os.Getenv("PATH"))
	// 变量os.Args是一个字符串slice
	// os.Args[0]是命令本身的名字
	// 实际参数范围是os.Args[1:len(os.Args)]，或写作os.Args[1:]
	fmt.Printf("argsFunc os.Args : %+v\n", os.Args)
}
