package predator

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coral struct {
	X, Y   float64
	Radius float64
	Status int
}

func NewCoral() *Coral {
	return &Coral{
		X:      rand.Float64() * Width,
		Y:      rand.Float64() * Height,
		Radius: 10,
	}
}

const (
	GrowthRate     = 1
	CoralMaxRadius = 50
)

const (
	CoralAlive int = iota
	CoralDead
)

func (c *Coral) Update() {
	if c.Status == CoralAlive {
		if c.Radius < 0 {
			c.Status = CoralDead
		} else if c.Radius < CoralMaxRadius {
			c.Radius += GrowthRate
		}
	}
}

func (c *Coral) Draw(screen *ebiten.Image) {
	DrawCircle(screen, c.X, c.Y, c.Radius, color.White)
}
