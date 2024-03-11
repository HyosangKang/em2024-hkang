package main

import "fmt"

func main() {
	fmt.Printf("%v\n", ExpBySequence(2.0, 10000000))
}

func Nadic(a, n int) []int {
	var l []int
	for a >= n {
		l = append(l, a%n) // a%n is the remainder of a / n
		a = a / n
	}
	l = append(l, a)
	return l
}

func ExpBySequence(x float64, n int) float64 {
	var f float64 = 1 + x/float64(n)
	var r float64 = 1
	var l []int = Nadic(n, 2)
	for i, k := range l { // i is the index and k is the value in the slice
		if k == 1 {
			var s float64 = f
			for j := 0; j < i; j++ {
				s = s * s
			}
			r = r * s
		}
	}
	return r
}

func ExpBySeries(x float64, n int) float64 {
	return 0
}
