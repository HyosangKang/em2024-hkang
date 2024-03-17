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
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			screen.Set(i, j, color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
