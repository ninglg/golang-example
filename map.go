package main

import (
	"fmt"
)

func main() {
	i := make(map[string]int)
	i["one"] = 1
	i["two"] = 2
	i["three"] = 3
	fmt.Println(i)
	fmt.Println(len(i))

	fmt.Println(i["one"])

	//
	delete(i, "three")
	fmt.Println(i)

	//
	a, b := i["two"]
	fmt.Println(a, b)

}
