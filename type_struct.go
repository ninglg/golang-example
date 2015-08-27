package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	x := new(Person)
	x.name = "XiaoMing"
	x.age = 17

	fmt.Println(x)
	fmt.Println(x.name)
}
