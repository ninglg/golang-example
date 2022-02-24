package main

import (
	"fmt"
	"sort"
)

func main() {
	student := make(map[string]int)
	student["aaa"] = 15
	student["bbb"] = 16
	student["ccc"] = 17

	for name := range student {
		fmt.Println(name)
	}

	for name, age := range student {
		fmt.Println("name = ", name, ", age = ", age)
	}

	age, ok := student["ccc"]
	if ok {
		fmt.Println("age = ", age)
	}

	delete(student, "ccc")
	age, ok = student["ccc"]
	if !ok {
		fmt.Println("age = nil")
	}

	// 对map的key进行排序
	var names []string
	for name := range student {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, student[name])
	}

	// map的值类型本身可以是符合数据类型，例如map或slice。
	// map的key必须是可比较类型，不可以是slice、map、function
}

// 和slice一样，map不可比较，唯一合法的比较就是和nil作比较。
