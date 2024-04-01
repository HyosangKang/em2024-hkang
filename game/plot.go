package game

import (
	"em2024-hkang/cal"

	"github.com/hajimehoshi/ebiten/v2"
)

// Requires ebiten
// Installation: type the following command in the terminal
// xcode-select --install (macos only)
// go get github.com/hajimehoshi/ebiten/v2

type Plot struct {
	F      cal.Function
	Xb, Yb [2]float64
	W, H   int
	N      int
}

func (g *Plot) Update() error {
	return nil
}

func (g *Plot) Draw(screen *ebiten.Image) {
	xys := g.F.Graph(g.N)
	for i := 0; i < g.N; i++ {
		i0, j0 := g.Pixel(xys[i][0], xys[i][1])
		i1, j1 := g.Pixel(xys[i+1][0], xys[i+1][1])
		Line(i0, j0, i1, j1, screen)
	}
}

func (g *Plot) Layout(_, _ int) (int, int) {
	return g.W, g.H
}
