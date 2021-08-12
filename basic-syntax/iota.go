package main

import "fmt"

const (
	aa = iota // 0
	bb        // 1
	cc        // 2
	dd = "ha" // 独立值"ha", iota += 1
	ee        // "ha"   iota += 1
	ff = 100  // iota +=1
	gg        // 100  iota +=1
	hh = iota // 7, 恢复计数
	ii        // 8
)

func main() {
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, ii)
	//0 1 2 ha ha 100 100 7 8
}
