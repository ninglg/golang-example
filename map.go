package main

import "fmt"

func mapFunc() {
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
}
