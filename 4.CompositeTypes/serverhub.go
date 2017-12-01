package main

import (
	"fmt"
)

type Server struct {
	ip   string
	port string
}

type ServerList []Server

func main() {
	var s ServerList
	s = append(s, Server{"127.0.0.1", "8080"})
	s = append(s, Server{"127.0.0.1", "9090"})

	for _, host := range s {
		fmt.Println(host.ip, ":", host.port)
	}
}
