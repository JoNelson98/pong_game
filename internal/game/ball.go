package game

type Ball struct {
	X, Y, W, H int
	DX, DY     int
}

func NewBall() Ball {
	return Ball{
		X: 0, Y: 200, W: 15, H: 15,
		DX: 3, DY: 3,
	}
}

func (b *Ball) Move() {
	b.X += b.DX
	b.Y += b.DY
}
