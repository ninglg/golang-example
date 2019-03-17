package main

import "fmt"

func stringFunc() {
	str := "hello, world!"
	fmt.Println("string : ", str)
	fmt.Println("string[2] : ", str[2])
	fmt.Println("string len : ", len(str))

	str2 := " good!"
	fmt.Println("string str+str2 : ", str+str2)

	for _, j := range str {
		fmt.Printf("%s ", string(j))
	}
	fmt.Println("")
}
