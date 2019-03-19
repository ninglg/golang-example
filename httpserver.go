package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func httpServerFunc() {
	http.HandleFunc("/hi", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
