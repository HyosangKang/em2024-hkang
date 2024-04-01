package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func Dot(x, y int, screen *ebiten.Image) {
	for i := -1; i <= 1; i++ {
		for j := -2; j <= 2; j++ {
			screen.Set(x+i, y+j, color.White)
			screen.Set(x+j, y+i, color.White)
		}
	}
}

func Line(x0, y0, x1, y1 int, screen *ebiten.Image) {
	dx := x1 - x0
	dy := y1 - y0
	if dx == 0 && dy == 0 {
		screen.Set(x0, y0, color.White)
		return
	}
	n := max(abs(dx), abs(dy))
	for i := 0; i <= n; i++ {
		x := int(float64(x0) + float64(i)*float64(dx)/float64(n))
		y := int(float64(y0) + float64(i)*float64(dy)/float64(n))
		screen.Set(x, y, color.White)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
