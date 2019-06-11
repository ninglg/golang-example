package main

import (
	"encoding/json"
	"fmt"
)

// srv struct Marshal
type srv struct {
	Name string
	IP   string
}

type srvSlice struct {
	SS []srv
}

// Animal struct Unmarshal
type Animal struct {
	Name  string
	Order string
	Price int  `json:"price, omitempty"`
}

func jsonFunc() {
	fmt.Println("========== jsonFunc ==========")
	var s srvSlice
	s.SS = append(s.SS, srv{Name: "Shanghai_VPN", IP: "127.0.0.1"})
	s.SS = append(s.SS, srv{Name: "Beijing_VPN", IP: "127.0.0.2"})
	// Marshal
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json Marshal err : ", err)
	}
	fmt.Println("json Marshal : ", string(b))

	// Unmarshal
	var jsonBlob = []byte(`[ 
		{"Name": "xiaoming", "Order": "abcde"}, 
		{"Name": "xiaogang", "Order": "fghij"}
		]`)

	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("json Unmarshal error : ", err)
	}
	fmt.Printf("json Unmarshal : %+v", animals)
	fmt.Println("")
}
