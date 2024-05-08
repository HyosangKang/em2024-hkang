package predator

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fish struct {
	X, Y   float64
	Prey   *Coral
	Status int
	Timer  int
}

func NewFish() *Fish {
	return &Fish{
		X: rand.Float64() * Width,
		Y: rand.Float64() * Height,
	}
}

const (
	FishSpeed       = 1
	FishRadius      = 5
	FishEatingSpeed = .005
)

const (
	FishHunting int = iota
	FishEating
	FishAlive
	FishDead
)

func (f *Fish) Update() {
	switch f.Status {
	case FishHunting:
		if f.Prey == nil {
			f.X += FishSpeed * (rand.Float64()*2 - 1)
			f.Y += FishSpeed * (rand.Float64()*2 - 1)
		} else {
			dx := f.Prey.X - f.X
			dy := f.Prey.Y - f.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < f.Prey.Radius {
				f.Status = FishEating
			}
		}
	case FishEating:
		f.Timer++
		f.Prey.Radius -= FishEatingSpeed
	}
}

func (f *Fish) Draw(screen *ebiten.Image) {
	DrawCircle(screen, f.X, f.Y, FishRadius, color.RGBA{60, 179, 113, 255})
}
