package cal

func (f Function) Graph(n int) [][2]float64 {
	var xys [][2]float64
	xarr := LinspaceFloat64(f.Domain[0], f.Domain[1], n)
	for _, x := range xarr {
		y := f.Evaluation(x)
		xys = append(xys, [2]float64{x, y})
	}
	return xys
}

func LinspaceFloat64(a, b float64, n int) []float64 {
	var arr []float64
	for i := 0; i <= n; i++ {
		arr = append(arr, a+(b-a)/float64(n)*float64(i))
	}
	return arr
}
