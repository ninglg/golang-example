package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// bufio可以简便的处理输入和输出
	// Scanner扫描器类型，可以读取输入，以行或单词为单位断开，这是处理以行为单位的输入内容的最简单方式
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// 注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		// if的条件不必放在圆括号中
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
