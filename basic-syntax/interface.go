package main

import "fmt"

type phone interface {
	call() string
}

type nokiaPhone struct {
}

func (np nokiaPhone) call() string {
	return "I am Nokia, I can call you!"
}

type iPhone struct {
}

func (ip iPhone) call() string {
	return "I am iPhone, I can call you!"
}

func interfaceFunc() {
	var ph phone

	ph = new(nokiaPhone)
	fmt.Println(ph.call())

	ph = new(iPhone)
	fmt.Println(ph.call())
}

