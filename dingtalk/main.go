package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var url string = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxx"

// SendMessage 发送钉钉机器人消息
func SendMessage(url, message string, ats ...string) (respContent string, err error) {
	c := &http.Client{
		Timeout: 1 * time.Second,
	}

	data := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": message},
	}

	if len(ats) != 0 {
		isAtAll := false
		var atMobiles []string
		for i := range ats {
			if ats[i] == "all" {
				isAtAll = true
			} else {
				atMobiles = append(atMobiles, ats[i])
			}
		}

		data["at"] = map[string]interface{}{
			"isAtAll":   isAtAll,
			"atMobiles": atMobiles,
		}
	}

	b, _ := json.Marshal(data)

	resp, err := c.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return "post请求失败", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		return string(body), nil
	}

	return string(body), errors.New(strconv.Itoa(resp.StatusCode))
}

func main() {
	result, err := SendMessage(url, "dingtalk测试消息通知", "all")
	if err != nil {
		fmt.Println("发送失败", result)
		return
	}

	fmt.Println("发送成功", result)
}