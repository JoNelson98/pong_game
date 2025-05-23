package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	Paddle    Paddle
	Ball      Ball
	Score     int
	HighScore int
}

func NewGame() *Game {
	return &Game{
		Paddle: NewPaddle(),
		Ball:   NewBall(),
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw paddle
	vector.DrawFilledRect(screen,
		float32(g.Paddle.X), float32(g.Paddle.Y),
		float32(g.Paddle.W), float32(g.Paddle.H),
		color.White, false,
	)
	// draw ball
	vector.DrawFilledRect(screen,
		float32(g.Ball.X), float32(g.Ball.Y),
		float32(g.Ball.W), float32(g.Ball.H),
		color.White, false,
	)
	// draw score
	scoreStr := fmt.Sprintf("Score: %d", g.Score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.White)

	highScoreStr := fmt.Sprintf("High Score: %d", g.HighScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 40, color.White)
}

func (g *Game) Update() error {
	g.Paddle.MoveOnKeyPress()
	g.Ball.Move()

	g.CollideWithWall()
	g.CollideWithPaddle()

	return nil
}

func (g *Game) Layout(w, h int) (int, int) {
	return 640, 400
}
