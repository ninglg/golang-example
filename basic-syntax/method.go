package main

import "fmt"

// 除了interface和pointer以外，任何类型都可以定义方法

// 方法不支持重载
// 方法的接收者名字没有限制，按惯例会选用简短有意义的名称（不推荐使用this、self）
// 如果方法内部并不引用实例，可省略参数名，仅保留类型
func main() {
	fmt.Println("method")
}

// 如何选择方法的receiver类型：
// 要修改实例状态，用*T
// 无须修改状态的小对象或固定值，建议用T
// 大对象建议用*T，以减少复制成本
// 引用类型、字符串、字典、函数等指针包装对象，直接用T
// 若包含Mutex等同步字段，用*T，避免因赋值造成锁操作无效
// 其他无法确定的情况都用*T
