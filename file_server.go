package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("/Users/ftp"))
	http.ListenAndServe(":8888", h)
}
