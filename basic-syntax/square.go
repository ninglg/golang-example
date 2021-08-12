package main

import (
	"fmt"
	"math"
)

func squareFunc() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
