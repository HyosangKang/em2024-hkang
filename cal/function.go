package cal

type Function struct {
	Domain     [2]float64
	Evaluation func(float64) float64
}

func (f Function) LeftSum(n int) float64 {
	var g float64 // g is going to be the integral of f on the domain
	a, b := f.Domain[0], f.Domain[1]
	for i := 0; i < n; i++ {
		x := a + (b-a)/float64(n)*float64(i)
		g += f.Evaluation(x) * (b - a) / float64(n)
	}
	return g
}

func (f Function) Integrate() float64 {
	var n int = 1000
	a, b := f.Domain[0], f.Domain[1]
	var g float64
	for i := 0; i <= n; i++ {
		x := a + (b-a)/float64(n)*float64(i)
		if i == 0 || i == n {
			g += f.Evaluation(x) / 2 * (b - a) / float64(n)
		} else {
			g += f.Evaluation(x) * (b - a) / float64(n)
		}
	}
	return g
}
