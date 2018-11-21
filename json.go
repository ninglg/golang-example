package main

import (
	"encoding/json"
	"fmt"
)

type Srv struct {
	SrvName string
	SrvIP   string
}

type SrvSlice struct {
	SrvS []Srv
}

func main() {
	var s SrvSlice
	s.SrvS = append(s.SrvS, Srv{SrvName: "Shanghai_VPN", SrvIP: "127.0.0.1"})
	s.SrvS = append(s.SrvS, Srv{SrvName: "Beijing_VPN", SrvIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
