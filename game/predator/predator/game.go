package predator

type Game struct {
	Humans []*Human
	Fishes []*Fish
	Corals []*Coral
}

const (
	Width  = 800
	Height = 600
)

func (g *Game) Layout(int, int) (int, int) {
	return Width, Height
}
