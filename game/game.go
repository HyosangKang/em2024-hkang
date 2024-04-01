package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Stage              int // 0: Ready 1: Throw 2: Land
	Width, Height      int
	Origin, CursorPos  [2]int
	Velocity, Position [2]float64
}

func (game *Game) StageReady() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		game.Velocity = [2]float64{float64(x-game.Origin[0]) / 50, float64(y-game.Origin[1]) / 50}
		game.CursorPos = [2]int{x, y}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		game.Stage = 1
	}
}

func (game *Game) StageThrow() {
	for i := 0; i < 2; i++ {
		game.Position[i] += game.Velocity[i]
	}
	if int(game.Position[0]) > game.Width || int(game.Position[1]) > game.Origin[1] {
		game.Stage = 2
	}
	game.Velocity[1] += 0.1
}

func (game *Game) StageLand() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		game.Position = [2]float64{float64(game.Origin[0]), float64(game.Origin[1])}
		game.Stage = 0
	}
}

func (game *Game) Update() error {
	switch game.Stage {
	case 0:
		game.StageReady()
	case 1:
		game.StageThrow()
	case 2:
		game.StageLand()
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	x, y := int(game.Position[0]), int(game.Position[1])
	switch game.Stage {
	case 0:
		Dot(game.Origin[0], game.Origin[1], screen)
		Line(game.Origin[0], game.Origin[1], game.CursorPos[0], game.CursorPos[1], screen)
	default:
		Dot(x, y, screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.Width, g.Height
}
