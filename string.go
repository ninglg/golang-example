package main

import (
	"fmt"
)

func main() {
	str := "hello, world!"
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(str[2])
	fmt.Printf("%c\n", str[2])
	fmt.Println(string(str[2]))

	str2 := " good!"
	fmt.Println(str + str2)

	str3 := "你好"
	fmt.Println(len(str3))

	for i, j := range str3 {
		fmt.Println(i, string(j))
	}
}