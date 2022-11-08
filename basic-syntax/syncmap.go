package main

import (
	"fmt"
	"sync"
)

// sync.Map不能通过make创建。无须初始化，直接声明即可
var sm sync.Map

func main() {
	// sync.Map 将键和值以 interface{} 类型进行保存
	sm.Store("aaa", 111)
	sm.Store("bbb", 222)
	sm.Store("ccc", 333)

	fmt.Println(sm.Load("aaa"))	// 111 true
	fmt.Println(sm.Load("ddd"))	// nil false

	sm.Delete("ccc")
	sm.Delete("eee")

	// Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回
	// Range() 参数中回调函数的返回值在需要继续迭代遍历时返回 true，终止迭代遍历时返回 false
	sm.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	// sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能
}
