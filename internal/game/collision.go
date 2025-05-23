package game

func (g *Game) CollideWithWall() {
	if g.Ball.X >= 640 {
		g.Reset()
	} else if g.Ball.X <= 0 {
		g.Ball.DX = 3
	}
	if g.Ball.Y <= 0 {
		g.Ball.DY = 3
	} else if g.Ball.Y >= 400 {
		g.Ball.DY = -3
	}
}

func (g *Game) CollideWithPaddle() {
	if g.Ball.X >= g.Paddle.X &&
		g.Ball.Y >= g.Paddle.Y &&
		g.Ball.Y <= g.Paddle.Y+g.Paddle.H {
		g.Ball.DX = -g.Ball.DX
		g.Score++
		if g.Score > g.HighScore {
			g.HighScore = g.Score
		}
	}
}

func (g *Game) Reset() {
	g.Ball = NewBall()
	g.Score = 0
}
