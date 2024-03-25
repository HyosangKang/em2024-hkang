package game

func (g *Game) Pixel(x, y float64) (int, int) {
	var i, j int
	i = int(5 * x / 4)
	j = int(7 - 7*y/3)
	return i, j
}
