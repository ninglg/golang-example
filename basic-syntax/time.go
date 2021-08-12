package main

import (
	"fmt"
	"time"
)

func timeFunc() {
	fmt.Println("========== timeFunc ==========")

	fmt.Println("timeFunc : ", time.Now().Format("2006-01-02 15:04:05"))
}
