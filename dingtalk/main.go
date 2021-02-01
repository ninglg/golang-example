package main

import (
	"fmt"
)

func main() {
	result, err := SendMessage("https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxx", "test测试消息通知", "all")
	if err != nil {
		fmt.Println("发送失败", result)
		return
	}

	fmt.Println("发送成功", result)
}