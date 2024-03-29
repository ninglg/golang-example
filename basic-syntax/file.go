package main

import (
	"fmt"
	"os"
)

func main() {
	// 将字符串写入文件
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, _ := f.WriteString("Hello World\n")

	fmt.Println(l, "bytes written successfully")

	// 将字节写入文件
	d := []byte{104, 101, 108, 108, 111}
	n, _ := f.Write(d)
	fmt.Println(n, "bytes written successfully")

	fmt.Fprintln(f)

	// 将字符串一行一行的写入文件
	d2 := []string{"welcome to the world", "it is beautiful"}
	for _, v := range d2 {
		_, _ = fmt.Fprintln(f, v)
	}

	_ = f.Close()
	// 追加到文件
	f, _ = os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	newLine := "File handling is easy."
	_, _ = fmt.Fprintln(f, newLine)
	fmt.Println("file appended successfully")

	_ = f.Close()
	// 删除文件
	_ = os.Remove("test.txt")
	fmt.Println("file removed successfully")
}
