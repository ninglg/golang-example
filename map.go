package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age int
	}

	var pp map[string]Person
	pp = make(map[string]Person)

	pp["one"] = Person{"XiaoWang", 18}
	pp["two"] = Person{"XiaoZhang", 17}

	fmt.Println(pp)

	data, ok := pp["two"]
	if ok {
		fmt.Println(data)
	} else {
		fmt.Println("data not exist!")
	}

	delete(pp, "two")
	fmt.Println(pp)
}
