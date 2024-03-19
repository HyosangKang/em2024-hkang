package main

import (
	"em2024-hkang/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	g := game.Game{
		SquareSize: 10,
	}
	ebiten.RunGame(&g)
}
