package cal

func Nadic(a, n int) []int {
	var l []int
	for a >= n {
		l = append(l, a%n) // a%n is the remainder of a / n
		a = a / n
	}
	l = append(l, a)
	return l
}
