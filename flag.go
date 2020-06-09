package main

import (
	"flag"
	"fmt"
)

func flagFunc() {
	fmt.Println("========== flagFunc ==========")

	var i bool
	var s string
	flag.BoolVar(&i, "i", false, "show info")
	flag.StringVar(&s, "s", "default", "show string")
	flag.Parse()
	fmt.Println("show info : ", i)
	fmt.Println("show string : ", s)
}

func main() {
	flagFunc()
}
