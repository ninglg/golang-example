package main

import (
	"fmt"
	"sort"
)

func mapFunc() {
	fmt.Println("========== mapFunc ==========")
	student := make(map[string]int)
	student["xiaowang"] = 15
	student["xiaoming"] = 16
	student["xiaogang"] = 17

	for name := range student {
		fmt.Println(name)
	}

	for name, age := range student {
		fmt.Println("name = ", name, ", age = ", age)
	}

	age, ok := student["xiaogang"]
	if ok {
		fmt.Println("age = ", age)
	}

	delete(student, "xiaogang")
	age, ok = student["xiaogang"]
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
}
