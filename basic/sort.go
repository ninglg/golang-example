package main

import (
	"fmt"
	"sort"
)

func sortFunc() {
	fmt.Println("========== sortFunc ==========")
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println("unsorted: ", s)

	sort.Ints(s)
	fmt.Println("sorted from small to big: ", s)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("sorted from big to small: ", s)
}

func main() {
	sortFunc()
}
