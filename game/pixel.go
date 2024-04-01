package game

func (g *Plot) Pixel(x, y float64) (int, int) {
	i := int((x - g.Xb[0]) / (g.Xb[1] - g.Xb[0]) * float64(g.W))
	j := int((g.Yb[1] - y) / (g.Yb[1] - g.Yb[0]) * float64(g.H))
	return i, j
}
