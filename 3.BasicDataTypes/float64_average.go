package main

import (
	"fmt"
)

func average(param []float64) (result float64) {
	sum := 0.0

	switch len(param) {
	case 0:
		result = 0
	default:
		for _, num := range param {
			sum += num
		}
		result = sum / float64(len(param))
	}
	return
}

func main() {
	input := []float64{1.2, 2.4, 3.6, 9.8}
	fmt.Println(average(input))
}
