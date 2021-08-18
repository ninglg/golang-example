package main

import (
	"fmt"
	"time"
)

//ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发。
//timer定时器，是到固定时间后会执行一次。
//如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool
func nTicker() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	i := 1
	for {
		select {
		case <- t.C:
			fmt.Printf("%d seconds ticker\n", 2*i)
			i++
		}
	}
}

func nTimer() {
	t := time.NewTimer(3 * time.Second)
	defer t.Stop()

	i := 1
	for {
		select {
		case <- t.C:
			fmt.Printf("%d seconds timer\n", 3*i)
			i++
			t.Reset(3 * time.Second)
		}
	}
}

func main() {
	go nTicker()
	go nTimer()
	time.Sleep(10 * time.Second)
}

