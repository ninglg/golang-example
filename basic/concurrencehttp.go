package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func concurrenceHTTP() {
	fmt.Println("========== concurrenceHttp ==========")
	urls := []string{"http://www.baidu.com", "http://www.qq.com", "http://www.taobao.com"}

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // 启动一个goroutine
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func main() {
	concurrenceHTTP()
}