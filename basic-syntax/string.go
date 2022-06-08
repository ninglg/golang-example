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
	fmt.Println("str4 SplitN : ", strings.SplitN(str4, "_", 3))
	fmt.Println("str4 SplitAfter : ", strings.SplitAfter(str4, "_"))
	fmt.Println("str4 SplitAfterN : ", strings.SplitAfterN(str4, "_", 3))

	str5 := `abc`
	s5 := []byte(str5)
	fmt.Println(string(s5))

	fmt.Println(strings.ToUpper("toupper"))
	fmt.Println(strings.ToLower("TOLOWER"))

	str6 := "你好，中国"
	fmt.Println(len(str6))                    //15
	fmt.Println(utf8.RuneCountInString(str6)) //5

	// 用加号连接两个字符串
	str7 := "this is " + "str7."
	fmt.Println(str7)

	// 用反引号表示多行字符串
	str8 := `this is multi-line
	string str8.
	`
	fmt.Println(str8)

	// 字符串虽然不能更改，但可以进行切片操作
	str9 := str8[5:]
	fmt.Println(str9)
}
