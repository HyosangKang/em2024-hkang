package virus

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	for _, a := range Agents {
		a.Draw(screen)
	}
}
