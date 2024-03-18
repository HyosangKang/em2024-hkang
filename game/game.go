package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Requires ebiten
// Installation: type the following command in the terminal
// xcode-select --install (macos only)
// go get github.com/hajimehoshi/ebiten/v2

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Set(320, 240, color.White)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 640, 480
}
