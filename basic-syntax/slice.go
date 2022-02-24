package main

import "fmt"

func main() {
	// make
	se := make([]int, 10, 20)

	// len
	fmt.Println("len(se) = ", len(se))
	// cap
	fmt.Println("cap(se) = ", cap(se))
	// fmt
	fmt.Println(se)

	se = []int{5, 6, 7, 8}
	fmt.Println("se = ", se)

	// append 1
	se = append(se, 9)
	fmt.Println("append 1 : ", se)
	se[1] = 9
	fmt.Println("slice : ", se[1:3])

	// append 2
	su := []int{1, 2, 3}
	sv := []int{4, 5, 6}
	su = append(su, sv...)
	fmt.Println("append 2 :", su)

	// copy
	st := []int{12, 34, 56, 78}
	sk := make([]int, 2)
	copy(sk, st)
	fmt.Println("slice copy : ", sk)
	// 部分cooy，比如移除其中一个元素。另外如果对顺序没要求，也可以把最后一个元素覆盖过来。
	copy(st[1:], st[2:])
	fmt.Println("slice part copy: ", st[:len(st)-1])

	// slice init
	// 给部分下标的位置进行赋值
	i := []int{3: 4, 6: 9}
	fmt.Println("slice2 : ", i)

	// slice range
	si := []float64{1.2, 2.4, 3.6, 9.8}
	fmt.Printf("slice range : ")
	for _, num := range si {
		fmt.Printf("%f ", num)
	}
}

// 和数组不同，slice无法做比较，不能用==来测试两个slice是否拥有相同的元素。但可以使用内置的比较函数。
// slice唯一允许的比较操作就是和nil做比较。
