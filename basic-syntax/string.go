package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
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

	str3 := []string{"hello", "test", "world"}
	fmt.Println("str3 : ", strings.Join(str3, "_"))

	str4 := "this_is_a_test"
	fmt.Println("str4 : ", strings.Split(str4, "_"))

	str5 := `abc`
	s5 := []byte(str5)
	fmt.Println(string(s5))

	fmt.Println(strings.ToUpper("toupper"))
	fmt.Println(strings.ToLower("TOLOWER"))

	str6 := "你好，中国"
	fmt.Println(len(str6))                    //15
	fmt.Println(utf8.RuneCountInString(str6)) //5
}
