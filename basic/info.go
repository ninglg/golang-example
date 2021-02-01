package main

import (
	"fmt"
	"unsafe"
)

// 常量
const (
	Info     = "abc"
	InfoLen  = len(Info)
	InfoSize = unsafe.Sizeof(Info)
)

func infoFunc() {
	fmt.Println(Info, InfoLen, InfoSize)
}

func main() {
	infoFunc()
}
