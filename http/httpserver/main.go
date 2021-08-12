package main

import (
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	_, _ = fmt.Fprintf(w, `{"code":200, "msg":"success", "data":"Hello World"}`)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	_ = http.ListenAndServe(":9999", nil)
}
