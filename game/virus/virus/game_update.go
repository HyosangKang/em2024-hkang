package virus

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		Agents = append(Agents, NewAgent())
	}
	for _, a := range Agents {
		a.Update(Agents)
	}
	return nil
}
