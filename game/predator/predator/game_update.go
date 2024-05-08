package predator

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.HandleInput()
	for _, c := range g.Corals {
		c.Update()
	}
	for _, f := range g.Fishes {
		if f.Status == FishAlive {
			f.Prey = g.FindCoral()
		}
		f.Update()
	}
	for _, h := range g.Humans {
		if h.Status == HumanHunting {
			h.Prey = g.FindFish()
		}
		h.Update()
	}
	return nil
}

func (g *Game) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.Humans = append(g.Humans, NewHuman())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.Fishes = append(g.Fishes, NewFish())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.Corals = append(g.Corals, NewCoral())
	}
}

func (g *Game) FindCoral() *Coral {
	var alive []*Coral
	for _, c := range g.Corals {
		if c.Status == CoralAlive {
			alive = append(alive, c)
		}
	}
	if len(alive) == 0 {
		return nil
	}
	return alive[rand.Intn(len(alive))]
}

func (g *Game) FindFish() *Fish {
	var alive []*Fish
	for _, f := range g.Fishes {
		if f.Status == FishAlive {
			alive = append(alive, f)
		}
	}
	if len(alive) == 0 {
		return nil
	}
	return alive[rand.Intn(len(alive))]
}
