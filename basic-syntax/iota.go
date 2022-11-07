package main

import "fmt"

const (
	aa = iota // 0
	bb        // 1
	cc        // 2
	dd = "ha" // 独立值"ha", iota += 1
	ee        // "ha"   iota += 1
	ff = 100  // iota += 1
	gg        // 100  iota += 1
	hh = iota // 7, 恢复计数
	ii        // 8
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func main() {
	//0 1 2 ha ha 100 100 7 8
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, ii)

	// 1 2 4 3
	// 当在一个const组中仅有一个标示符在一行的时候，它将使用增长的iota取得前面的表达式并且再运用它
	// 在Go语言的spec中，这就是所谓的隐性重复最后一个非空的表达式列表
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
}
