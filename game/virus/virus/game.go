package virus

type Game struct{}

var (
	Agents []*Agent
)

const (
	InitNumAgents = 100
	InitExposed   = 20
)

func init() {
	for i := 0; i < InitNumAgents; i++ {
		Agents = append(Agents, NewAgent())
	}
	for i := 0; i < InitExposed; i++ {
		a := NewAgent()
		a.State = Exposed
		Agents = append(Agents, a)
	}
}

const (
	Width  = 800
	Height = 600
)

func (g *Game) Layout(int, int) (int, int) {
	return Width, Height
}
