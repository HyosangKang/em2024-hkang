package hw

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	HW    []*HW
	Seed  float64
	Tick  int
	Timer int
}

var Me *ME

func init() {
	Me = &ME{
		Y: Height - MeSize,
		S: 10,
	}
}

func (g *Game) Update() error {
	g.Tick++
	if g.Tick%10 == 0 {
		g.Seed += 0.025
		r := rand.Float64()
		if r < 0.2+g.Seed*0.01 {
			g.HW = append(g.HW, NewHW(g.Seed))
		}
	}
	for _, h := range g.HW {
		h.Update()
	}
	Me.Update()
	Me.Contact(g.HW)
	return nil
}

var (
	TitleFontFace *text.GoTextFaceSource

	//go:embed font.ttf
	font_ttf []byte
)

func init() {
	TitleFontFace, _ = text.NewGoTextFaceSource(bytes.NewReader(font_ttf))
}

func DrawText(screen *ebiten.Image, msg string, size float64, x, y float64) {
	text.Draw(screen, msg, &text.GoTextFace{
		Source: TitleFontFace,
		Size:   size,
	}, NewTextDrawOption(x, y))
}

func NewTextDrawOption(x, y float64) *text.DrawOptions {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.White)
	op.PrimaryAlign = text.AlignStart
	op.SecondaryAlign = text.AlignStart
	return op
}

func (g *Game) Draw(scr *ebiten.Image) {
	DrawText(scr, fmt.Sprintf("%d", g.Timer), 30, Width-60, 15)
	DrawText(scr, fmt.Sprintf("%d", Me.Score), 30, Width-100, 15)
	if g.Tick%60 == 0 {
		g.Tick = 0
		g.Timer++
	}
	for _, h := range g.HW {
		h.Draw(scr)
	}
	Me.Draw(scr)
}

const (
	Width  = 800
	Height = 600
)

func (g *Game) Layout(int, int) (int, int) {
	return Width, Height
}
