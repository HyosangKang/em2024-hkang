package cal

import "fmt"

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
	var r, f float64 = 1, 1 // r is the result and f is the i the term.
	for i := 1; i <= n; i++ {
		f *= x / float64(i)
		r += f
	}
	return r
}

var Eps = 1e-10

func IsZero(f float64) bool {
	return -Eps < f && f < Eps
}

func Exp(x float64) float64 {
	var r, f float64 = 1, 1 // r is the result and f is the i the term.
	for i := 1; !IsZero(f); i++ {
		fmt.Println(f)
		f *= x / float64(i)
		r += f
	}
	return r
}
