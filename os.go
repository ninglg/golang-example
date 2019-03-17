package main

import (
	"fmt"
	"os"
)

func osFunc() {
	fmt.Println("osFunc : ", os.Getenv("GOPATH"))
}
