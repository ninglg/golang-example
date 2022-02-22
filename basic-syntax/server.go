package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8092", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%s", r.Form.Get("a"))
}
