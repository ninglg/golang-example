package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func structFunc() {
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
	z := person{"XiaoHong", 19}
	fmt.Println("struct z : ", z)

	// 4
	g := person{age: 17, name: "XiaoZhang"}
	fmt.Println("struct g : ", g)
}
