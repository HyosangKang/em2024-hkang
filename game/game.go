package game

import (
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
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.X, g.Y = ebiten.CursorPosition()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	Dot(g.X, g.Y, screen)
}

func (*Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 640, 480
}
