package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 480
	birdSize     = 32
	gravity      = 0.81
	jumpSpeed    = -10
	gapSize      = 120
	pipeWidth    = 60
	pipeSpeed    = 2
	dt           = 1.0
)

type Game struct {
	birdY           float64
	birdVelocity    float64
	pipeX           float64
	upperPipeHeight int
	score           int
	gameOver        bool
	passedPipe      bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset()
		}
		return nil
	}

	// 새의 위치&속도를 정하기 위한 Euler`s method
	g.birdVelocity += gravity * dt
	g.birdY += g.birdVelocity * dt

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.birdVelocity = jumpSpeed
	}

	g.pipeX -= pipeSpeed
	if g.pipeX < -pipeWidth {
		g.pipeX = screenWidth
		g.upperPipeHeight = rand.Intn(screenHeight - gapSize)
		g.passedPipe = false
	}

	if !g.passedPipe && g.pipeX < screenWidth/2-birdSize/2 {
		g.score++
		g.passedPipe = true
	}

	if g.birdY < 0 || g.birdY+birdSize > screenHeight || g.hitPipe() {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 135, G: 206, B: 235, A: 255}) // sky blue background

	// 장애물 구현(파이프 같이 생긴 장애물 말고 다른 모양으로 바뀔 예정)
	pipeColor := color.RGBA{R: 0, G: 128, B: 0, A: 255}
	ebitenutil.DrawRect(screen, g.pipeX, 0, pipeWidth, float64(g.upperPipeHeight), pipeColor)
	ebitenutil.DrawRect(screen, g.pipeX, float64(g.upperPipeHeight+gapSize), pipeWidth, screenHeight-float64(g.upperPipeHeight+gapSize), pipeColor)

	// 새 구현(새 말고 다른 모양으로 바뀔 예정)
	birdColor := color.RGBA{R: 255, G: 255, B: 0, A: 255} //일단 원작 게임에 맞춰 노란색
	ebitenutil.DrawRect(screen, screenWidth/2-birdSize/2, g.birdY, birdSize, birdSize, birdColor)

	// 게임 점수 표시
	scoreText := fmt.Sprintf("Score: %d", g.score)
	ebitenutil.DebugPrintAt(screen, scoreText, 10, 10)

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press Space to Restart")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) hitPipe() bool {
	return (screenWidth/2+birdSize/2 > g.pipeX && screenWidth/2-birdSize/2 < g.pipeX+pipeWidth) &&
		(g.birdY < float64(g.upperPipeHeight) || g.birdY+birdSize > float64(g.upperPipeHeight+gapSize))
}

func (g *Game) reset() {
	g.birdY = screenHeight / 2
	g.birdVelocity = 0
	g.pipeX = screenWidth
	g.upperPipeHeight = rand.Intn(screenHeight - gapSize)
	g.score = 0
	g.gameOver = false
	g.passedPipe = false
}

func main() {
	game := &Game{}
	game.reset()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Bird (Go with Ebiten)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
