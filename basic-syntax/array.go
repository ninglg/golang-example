package main

import "fmt"

func main() {
	arr1 := [3]int{22, 33, 44}
	arr2 := [...]string{"hello", "world"}
	// 定义多维数组时，仅第一维度允许使用3个点
	arr3 := [...][3]int{{4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// 如果一个数组的元素类型是可比较的，那么这个数组也是可比较的

// 也可以显示的传递一个数组的指针给函数。比如以下版本的数组清零程序：
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
