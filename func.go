package main

import "fmt"

func main() {
	test()
}

func sum(x int, y int) (s int) {
	return x + y
}

func p(x int, y int) (int, int) {
	return x, y
}

func test() {
	fmt.Println("Hello, function!")
	fmt.Println(sum(3, 4))
	fmt.Println(p(6, 7))

	a, b := p(8, 9)
	fmt.Println(b, a)
}
