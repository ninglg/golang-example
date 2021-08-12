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

	fmt.Println(r.Form)
	fmt.Println("path : ", r.URL.Path)
	fmt.Println("scheme : ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	_, _ = fmt.Fprintln(w, r.Form)
}

func main() {
	http.HandleFunc("/hi", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
