package main

import (
	"fmt"
)

func fibonacci2(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func show(c, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func fibFunc2() {
	data := make(chan int)
	leave := make(chan int)
	go show(data, leave)
	go fibonacci2(data, leave)
	//for {
	//	time.Sleep(1)
	//}
}
