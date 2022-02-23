package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	a, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

	k := strconv.FormatInt(int64(a), 10)
	fmt.Println(k)

}
