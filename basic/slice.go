package main

import "fmt"

func sliceFunc() {
	fmt.Println("========== sliceFunc ==========")
	// make
	se := make([]int, 10, 20)
	// len
	fmt.Println("len(se) = ", len(se))
	// cap
	fmt.Println("cap(se) = ", cap(se))
	se = []int{5, 6, 7, 8}
	fmt.Println("se = ", se)
	// append 1
	se = append(se, 9)
	fmt.Println("append 1 : ", se)
	se[1] = 9
	fmt.Println("slice : ", se[1:3])
	// append 2
	su := []int{1, 2, 3}
	sv := []int{4, 5, 6}
	su = append(su, sv...)
	fmt.Println("append 2 :", su)

	// copy
	st := []int{12, 34}
	sk := make([]int, 2)
	copy(sk, st)
	fmt.Println("slice copy : ", sk)

	// slice init
	i := []int{3: 4, 6: 9}
	fmt.Println("slice2 : ", i)

	// slice range
	si := []float64{1.2, 2.4, 3.6, 9.8}
	fmt.Printf("slice range : ")
	for _, num := range si {
		fmt.Printf("%f ", num)
	}
	fmt.Println("")
}
