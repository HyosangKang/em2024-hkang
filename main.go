package main

import "fmt"

func main() {
	var a float64 = 1
	var n int = 100

	fmt.Printf("%v\n", ExpBySequence(a, n))
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
	for i := 0; i < n; i++ {
		r = r * f
	}
	return r
}
