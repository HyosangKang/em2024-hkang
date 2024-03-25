package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Requires ebiten
// Installation: type the following command in the terminal
// xcode-select --install (macos only)
// go get github.com/hajimehoshi/ebiten/v2

type Game struct {
	X, Y       int
	SquareSize int
}

func (g *Game) Update() error {
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	g.X, g.Y = ebiten.CursorPosition()
	// }
	g.X, g.Y = g.Pixel(1, 1)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			Dot(i, j, screen)
		}
	}
	i, j := g.Pixel(0.5, 0.5)
	screen.Set(i, j, color.Black)
	i, j = g.Pixel(1, 1)
	screen.Set(i, j, color.Black)
	i, j = g.Pixel(2, 2)
	screen.Set(i, j, color.Black)
}

func (*Game) Layout(_, _ int) (int, int) {
	return 5, 7
}
