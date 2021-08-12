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

func main() {
	fmt.Println(Info, InfoLen, InfoSize)
}
