package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s ", s)
	}
}

func goFunc() {
	fmt.Println("========== goFunc ==========")
	go say("world")
	say("hello")
	fmt.Println("")
}

func main() {
	goFunc()
}