package main

import (
	"fmt"
	sc "strconv"
)

func main() {
	v := "10"
	if s, err := sc.Atoi(v); err == nil {
		fmt.Printf("importFunc : %T, %v\n", s, s)
	}
}
