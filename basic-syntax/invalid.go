package main

import "fmt"

func main() {
	i, j := 2, 6
	//j = i++  // 不支持
	//--j  // 不支持
	fmt.Println("invalidFunc : ", i, j)
}
