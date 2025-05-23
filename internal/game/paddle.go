package game

import (
	"github.com/JoNelson98/go_pong/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	X, Y, W, H int
}

func NewPaddle() Paddle {
	return Paddle{
		X: 600, Y: 200, W: 15, H: 100,
	}
}

func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyJ) && p.Y + p.H < config.ScreenHeight{
		p.Y += 6
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) && p.Y > 0 {
		p.Y -= 6
	}
}
