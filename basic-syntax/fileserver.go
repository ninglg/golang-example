package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/Users/ftp"))
	_ = http.ListenAndServe(":8888", fs)
}
