package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("timeFunc : ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("timeFunc : ", time.Now().Format("2006年1月2号 3:4:5"))
}
