package hw

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

var (
	MeImage *ebiten.Image
)

func init() {
	var err error
	MeImage, _, err = ebitenutil.NewImageFromFile("hw/me.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ME) Draw(scr *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.5, .5)
	op.GeoM.Translate(m.X, m.Y-50)
	scr.DrawImage(MeImage, op)
	// for i := -m.S; i <= m.S; i++ {
	// 	for j := -m.S; j <= m.S; j++ {
	// 		scr.Set(int(m.X)+i, int(m.Y)+j, color.White)
	// 	}
	// }
}
