package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func httpGet() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	var body, err2 = ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("https://www.baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("https://www.baidu.com",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func httpDo() {
	client := &http.Client{
		Timeout: time.Second,
	}

	req, err := http.NewRequest("POST", "https://www.baidu.com", strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=test")

	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main() {
	go httpGet()
	time.Sleep(time.Second)

	go httpPost()
	time.Sleep(time.Second)

	go httpPostForm()
	time.Sleep(time.Second)

	go httpDo()
	time.Sleep(time.Second)
}
