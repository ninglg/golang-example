package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	x := new(person)
	x.name = "XiaoMing"
	x.age = 17

	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(x.name)

	var y person
	y.name = "XiaoGang"
	y.age = 23
	fmt.Println(y)
	fmt.Println(y.name)

	z := person{"XiaoHong", 19}
	fmt.Println(z)

	g := person{age: 17, name: "XiaoZhang"}
	fmt.Println(g)
}
