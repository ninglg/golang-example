package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello, world!\n")
	_ = r.ParseForm()

	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, "r.URL.Path: "+r.URL.Path)
	fmt.Fprintln(w, "r.URL.Scheme: "+r.URL.Scheme)
	fmt.Fprintln(w, r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Fprintln(w, "key:", k)
		fmt.Fprintln(w, "val:", strings.Join(v, ""))
	}

	_, _ = fmt.Fprintln(w, r.Form)
}

func main() {
	http.HandleFunc("/hi", hello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

