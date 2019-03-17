package main

import "fmt"

func sliceFunc() {
	se := make([]int, 10, 20)
	fmt.Println("slice len(se) = ", len(se))
	fmt.Println("slice cap(se) = ", cap(se))
	se = []int{5, 6, 7, 8}
	fmt.Println("slice se = ", se)
	se = append(se, 9)
	fmt.Println("slice append se = ", se)
	fmt.Println("slice : ", se[1:3])

	//
	si := []float64{1.2, 2.4, 3.6, 9.8}
	fmt.Printf("slice range : ")
	for _, num := range si {
		fmt.Printf("%f ", num)
	}
	fmt.Println("")
}
