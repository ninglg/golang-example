package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	str := "init"
	if len(str) > 0 {
		fmt.Println(str)
	}

	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(time.Now())
}
