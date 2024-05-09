package predator

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.HandleInput()
	for _, c := range g.Corals {
		c.Update()
	}
	for _, f := range g.Fishes {
		f.Update(g.Corals)
	}
	for _, h := range g.Humans {
		h.Update(g.Fishes)
	}
	return nil
}

func (g *Game) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		if len(g.Humans) < MaxNum {
			g.Humans = append(g.Humans, NewHuman())
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if len(g.Humans) < MaxNum {
			g.Fishes = append(g.Fishes, NewFish())
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if len(g.Humans) < MaxNum {
			g.Corals = append(g.Corals, NewCoral())
		}
	}
}
