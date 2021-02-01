package main

import (
	"fmt"
	sc "strconv"
)

func importFunc() {
	v := "10"
	if s, err := sc.Atoi(v); err == nil {
		fmt.Printf("importFunc : %T, %v", s, s)
	}
	fmt.Println("")
}

func main() {
	importFunc()
}
