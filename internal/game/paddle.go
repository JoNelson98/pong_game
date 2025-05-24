package game

import (
	"github.com/JoNelson98/go_pong/internal/config"
	utils "github.com/JoNelson98/go_pong/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	X, Y, W, H        int
	UpDuration        float64 // Duration K key is held
	DownDuration      float64 // Duration J key is held
	CurrentSpeed      float64 // Current speed of paddle movement
	MaxSpeed          float64 // Initial fast speed
	CruiseSpeed       float64 // Slower speed after easing
	AccelerationTime  float64 // Time to transition to cruise speed
}

func NewPaddle() Paddle {
	return Paddle{
		X:                600,
		Y:                200,
		W:                15,
		H:                100,
		MaxSpeed:         10.0, // Initial fast speed
		CruiseSpeed:      6.0,  // Slower cruise speed
		AccelerationTime: 0.5,  // Time to transition (in seconds)
	}
}

func (p *Paddle) MoveOnKeyPress() {
	// Time delta per frame (assuming 60 FPS)
	dt := 1.0 / 60.0

	// Handle K key (move up)
	if ebiten.IsKeyPressed(ebiten.KeyK) && p.Y > 0 {
		p.CurrentSpeed = utils.CalculateEasedSpeed(p.UpDuration, p.MaxSpeed, p.CruiseSpeed, p.AccelerationTime)
		p.Y -= int(p.CurrentSpeed)
		p.UpDuration += dt
		p.DownDuration = 0 // Reset opposite direction
	} else {
		p.UpDuration = 0 // Reset when key is released
	}

	// Handle J key (move down)
	if ebiten.IsKeyPressed(ebiten.KeyJ) && p.Y+p.H < config.ScreenHeight {
		p.CurrentSpeed = utils.CalculateEasedSpeed(p.DownDuration, p.MaxSpeed, p.CruiseSpeed, p.AccelerationTime)
		p.Y += int(p.CurrentSpeed)
		p.DownDuration += dt
		p.UpDuration = 0 // Reset opposite direction
	} else {
		p.DownDuration = 0 // Reset when key is released
	}

	// Ensure paddle stays within bounds
	if p.Y < 0 {
		p.Y = 0
	}
	if p.Y+p.H > config.ScreenHeight {
		p.Y = config.ScreenHeight - p.H
	}
}