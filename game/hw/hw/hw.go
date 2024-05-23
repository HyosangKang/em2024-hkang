package hw

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type HW struct {
	Vel  float64 // velocity
	X, Y float64 // position
	G    float64 // gravity
	S    int     // size
	Gone bool
}

func NewHW(seed float64) *HW {
	hw := &HW{}
	hw.S = rand.Intn(10) + 5
	hw.G = rand.Float64() * seed * 0.1
	hw.X = float64(hw.S) + rand.Float64()*float64(Width-2*hw.S)
	hw.Vel = rand.Float64() * seed
	return hw
}

func (h *HW) Update() {
	h.Y += h.Vel
	h.Vel += h.G
	if h.Y >= Height {
		h.Gone = true
	}
}

func (h *HW) Draw(scr *ebiten.Image) {
	if h.Gone {
		return
	}
	for i := -h.S; i <= h.S; i++ {
		for j := -h.S; j <= h.S; j++ {
			scr.Set(int(h.X)+i, int(h.Y)+j, color.White)
		}
	}
}
