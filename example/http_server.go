package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
