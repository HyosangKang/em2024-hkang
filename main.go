package main

import (
	"em2024-hkang/cal"
	"em2024-hkang/game"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println(cal.Exp(1))
	ebiten.SetWindowSize(600, 600)
	g := game.Game{
		F: cal.Function{
			Domain:     [2]float64{0, 1},
			Evaluation: cal.Exp,
		},
		W:  600,
		H:  600,
		Xb: [2]float64{-.1, 1.1},
		Yb: [2]float64{.9, 3.1},
		N:  20,
	}
	fmt.Println(cal.Exp(2))
	ebiten.RunGame(&g)
}
