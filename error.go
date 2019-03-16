package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type divideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *divideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := divideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	}

	return varDividee / varDivider, ""
}

func errorFunc() {
	// 正常情况
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
