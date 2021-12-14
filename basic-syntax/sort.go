package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println("unsorted: ", s)

	// 从小到大顺序排序
	sort.Ints(s)
	fmt.Println("sorted from small to big: ", s)

	// 从大到小逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("sorted from big to small: ", s)

	// 自定义排序方法
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))

}

func largestNumber(nums []int) string {
	flag := true
	var nums2 []string
	for _, i := range nums {
		nums2 = append(nums2, strconv.Itoa(i))
		if flag && (i != 0) {
			flag = false
		}
	}

	if flag {
		return "0"
	}

	sort.Sort(mySort(nums2))

	result := ""
	for _, i := range nums2 {
		result += i
	}

	return result
}

type mySort []string

func (s mySort) Len() int {
	return len(s)
}
func (s mySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s mySort) Less(i, j int) bool {
	if s[i]+s[j] > s[j]+s[i] {
		return true
	}

	return false
}
