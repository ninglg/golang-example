package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	p := []color.Color{color.Black, color.White}
	fmt.Println(p)
}
