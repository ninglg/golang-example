package main

import (
	"container/list"
	"fmt"
)

func main() {
	tmpList := list.New()

	for i := 1; i <= 10; i += 2 {
		tmpList.PushBack(i)
	}

	tmpList.PushFront(0)

	for j := tmpList.Front(); j != nil; j = j.Next() {
		fmt.Println(j.Value)
	}
}

