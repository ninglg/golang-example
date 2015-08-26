package main

import "fmt"

func main() {
	var arr [10]int
	brr := [5]int{6, 7, 8, 9, 10}
	crr := [...]int{3: 2, 4: 1}

	for i := 0; i < 10; i++ {
		arr[i] = i * 2
	}
	fmt.Println(arr)

	fmt.Println(brr)

	fmt.Println(crr)

	fmt.Println(len(crr))

}
