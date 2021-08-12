package main

import (
	"fmt"
)

func main() {
	var r interface{}

	r = "abc"

	switch r.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	}
}
