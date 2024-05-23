package hw

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ME struct {
	X, Y  float64 // position
	S     int     // size
	Score int
}

const (
	MeSize = 10
	Speed  = 4
)

func (m *ME) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		m.X += Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m.X -= Speed
	}
}

func (m *ME) Contact(hw []*HW) {
	for _, h := range hw {
		if h.Gone {
			continue
		}
		if math.Abs(m.X-h.X) < float64(m.S+h.S) && math.Abs(m.Y-h.Y) < float64(m.S+h.S) {
			m.Score += h.S
			h.Gone = true
		}
	}
}

func (m *ME) Draw(scr *ebiten.Image) {
	for i := -m.S; i <= m.S; i++ {
		for j := -m.S; j <= m.S; j++ {
			scr.Set(int(m.X)+i, int(m.Y)+j, color.White)
		}
	}
}
