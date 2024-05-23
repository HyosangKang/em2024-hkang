package main

import (
	"cannon/cannon"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(600, 200)
	g := cannon.Game{
		Width:    300,
		Height:   100,
		Stage:    0,
		Origin:   [2]int{20, 80},
		Position: [2]float64{20, 80},
	}
	ebiten.RunGame(&g)
}
