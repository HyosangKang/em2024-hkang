package predator

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	for _, c := range g.Corals {
		c.Draw(screen)
	}
	for _, f := range g.Fishes {
		f.Draw(screen)
	}
	for _, h := range g.Humans {
		h.Draw(screen)
	}
}

const (
	Nsub = 100
)

func DrawCircle(screen *ebiten.Image, x, y, r float64, color color.Color) {
	t := Linspace(0, 2*math.Pi, Nsub)
	for i := 0; i < Nsub; i++ {
		x0, y0 := x+r*math.Cos(t[i]), y+r*math.Sin(t[i])
		x1, y1 := x+r*math.Cos(t[(i+1)%Nsub]), y+r*math.Sin(t[(i+1)%Nsub])
		DrawLine(screen, x0, y0, x1, y1, color)
	}
}

func Linspace(a, b float64, n int) []float64 {
	if n < 2 {
		return nil
	}
	t := make([]float64, n)
	for i := 0; i < n; i++ {
		t[i] = a + (b-a)*float64(i)/float64(n-1)
	}
	return t
}

func DrawLine(screen *ebiten.Image, x0, y0, x1, y1 float64, color color.Color) {
	dx := x1 - x0
	dy := y1 - y0
	if dx == 0 && dy == 0 {
		screen.Set(int(x0), int(y0), color)
		return
	}
	n := AbsMax(dx, dy)
	for i := 0; i <= int(n); i++ {
		x := x0 + float64(i)*dx/n
		y := y0 + float64(i)*dy/n
		screen.Set(int(x), int(y), color)
	}
}

func AbsMax(x, y float64) float64 {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x < y {
		return y
	}
	return x
}

func Translate(x, y float64) (float64, float64) {
	return x + Width/2, y + Height/2
}
