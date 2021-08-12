package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("rand int = ", rand.Int())
	fmt.Println("rand float64 = ", rand.Float64())
}
