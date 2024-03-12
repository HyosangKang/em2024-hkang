package main

import (
	"em2024-hkang/cal"
	"fmt"
)

func main() {
	f := cal.Function{
		Domain: [2]float64{0.0, 1.0},
		Evaluation: func(x float64) float64 {
			return x * x
		},
	}
	fmt.Printf("%v\n", f.LeftSum(1000))
	fmt.Printf("%v\n", f.Integrate())
}
