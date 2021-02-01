package main

import (
	"fmt"
)

type address struct {
	addr string
	city string
}
type person struct {
	name string
	age  int
	friends []*person  // 可用指针类型创建递归数据结构，如链表或树。
	address
}

func structFunc() {
	fmt.Println("========== structFunc ==========")
	// 1
	x := new(person)
	x.name = "XiaoMing"
	x.age = 17
	fmt.Println("struct x : ", x)
	fmt.Println("struct *x : ", *x)
	fmt.Println("struct x.name : ", x.name)

	// 2
	var y person
	y.name = "XiaoGang"
	y.age = 23
	fmt.Println("struct y : ", y)
	fmt.Println("struct y.name : ", y.name)

	// 3
	z := person{"XiaoHong", 19, nil, address{}}
	fmt.Println("struct z : ", z)

	// 4
	g := person{age: 17, name: "XiaoZhang", friends:nil, address:address{}}
	fmt.Println("struct g : ", g)
}

func main() {
	structFunc()
}