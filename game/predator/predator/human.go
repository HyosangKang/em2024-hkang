package predator

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Human struct {
	X, Y   float64
	Status int
	Timer  int
	Prey   *Fish
}

func NewHuman() *Human {
	return &Human{
		X: rand.Float64() * Width,
		Y: rand.Float64() * Height,
	}
}

const (
	HumanEating int = iota
	HumanSleeping
	HumanHunting
)

const (
	SleepingTime = 240
	EatingTime   = 120
)

const (
	HumanRadius = 10
	HumanSpeed  = 3
)

func (h *Human) Update() {
	switch h.Status {
	case HumanEating:
		h.Timer++
		if h.Timer >= EatingTime {
			h.Status = HumanSleeping
			h.Timer = 0
		}
	case HumanSleeping:
		h.Timer++
		if h.Timer >= SleepingTime {
			h.Status = HumanHunting
			h.Timer = 0
		}
	case HumanHunting:
		dx := h.Prey.X - h.X
		dy := h.Prey.Y - h.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist < FishRadius {
			h.Prey.Status = FishDead
			h.Status = HumanEating
			h.Timer = 0
		} else {
			h.X += (h.Prey.X - h.X) / dist * HumanSpeed
			h.Y += (h.Prey.Y - h.Y) / dist * HumanSpeed
		}
	}
}

func (h *Human) Draw(screen *ebiten.Image) {
	DrawCircle(screen, h.X, h.Y, HumanRadius, color.RGBA{255, 165, 0, 255})
}
