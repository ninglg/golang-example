package main

import (
	"fmt"
)

type address struct {
	addr string
	city string
}

// 结构体可以嵌套
type person struct {
	name    string
	age     int
	friends []*person // 可用指针类型创建递归数据结构，如链表或树。
	address           // 匿名成员。匿名成员不一定是结构体类型，任何命名类型或者指向命名类型的指针都可以。
}

func main() {
	// 1
	x := new(person)
	x.name = "aaa"
	x.age = 17
	fmt.Println("struct x : ", x)
	fmt.Println("struct *x : ", *x)
	fmt.Println("struct x.name : ", x.name)

	// 2
	var y person
	y.name = "bbb"
	y.age = 23
	fmt.Println("struct y : ", y)
	fmt.Println("struct y.name : ", y.name)

	// 3
	z := person{"ccc", 19, nil, address{}}
	fmt.Println("struct z : ", z)

	// 4
	g := person{age: 17, name: "ddd", friends: nil, address: address{}}
	fmt.Println("struct g : ", g)
}

// 如果结构体的所有成员变量都可以比较，那么这个结构体就是可比较的。和其它可比较的类型一样，可比较的结构体都可以作为map的键类型。
