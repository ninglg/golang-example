package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout:       1000 * time.Millisecond,
	}

	result, err := client.Get("http://127.0.0.1:9999/")
	if err != nil {
		fmt.Printf("err: %+v\v", err)
		return
	}

	fmt.Printf("result: %+v\n", result)
}