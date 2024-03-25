package game

import (
	"em2024-hkang/cal"

	"github.com/hajimehoshi/ebiten/v2"
)

// Requires ebiten
// Installation: type the following command in the terminal
// xcode-select --install (macos only)
// go get github.com/hajimehoshi/ebiten/v2

type Game struct {
	F      cal.Function
	Xb, Yb [2]float64
	W, H   int
	N      int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	xys := g.F.Graph(g.N)
	for _, xy := range xys {
		i, j := g.Pixel(xy[0], xy[1])
		Dot(i, j, screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.W, g.H
}
