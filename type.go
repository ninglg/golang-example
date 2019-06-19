package main

import (
	"fmt"
)

func typeFunc() {
	fmt.Println("========== typeFunc ==========")
	var r interface{}
	r = "abc"
	switch r.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	}
}