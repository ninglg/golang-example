package main

import "fmt"

// 定义结构体
type Circle struct {
	radius float64
}

// 该method属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func circleFunc() {
	// 结构体
	c1 := Circle{10.00}
	fmt.Println("圆的半径 = ", c1.radius)

	// 方法
	var c2 Circle
	c2.radius = 10.00
	fmt.Println("圆的面积 = ", c2.getArea())
}
