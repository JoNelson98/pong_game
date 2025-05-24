package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type GameState int


const (
	StateStartScreen GameState = iota
	StateSinglePlayer
	StateMultiplayerMenu
	StateMultiplayer
	StateGameOver
)
type Game struct {
	State GameState
	Paddle    Paddle
	Ball      Ball
	Score     int
	HighScore int
	PrevKeys map[ebiten.Key]bool
	PlayerName string
}

func NewGame() *Game {
	return &Game{
		State: StateStartScreen,
		Paddle: NewPaddle(),
		Ball:   NewBall(),
		PrevKeys: make(map[ebiten.Key]bool),
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case StateStartScreen:
		text.Draw(screen,"PONG", basicfont.Face7x13,300,180, color.White)
		text.Draw(screen,"Press 1 for Single Player", basicfont.Face7x13,220,220, color.White)
		text.Draw(screen,"Press 2 for Multiplayer", basicfont.Face7x13,220,240, color.White)
	case StateSinglePlayer:
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
	case StateMultiplayerMenu:
			text.Draw(screen,"MULTIPLAYER MENU",basicfont.Face7x13,280,120,color.White )
			text.Draw(screen,"Press H to host",basicfont.Face7x13,180,140,color.White )
			text.Draw(screen,"Press J to Join",basicfont.Face7x13,180,160,color.White )
	case StateGameOver:
			text.Draw(screen,"GAME OVER", basicfont.Face7x13,280,120, color.White)
			text.Draw(screen,"Press 1 to restart Press 2 for main menu", basicfont.Face7x13,180,140, color.White)
	}
	
}

func (g *Game) Update() error {
	switch g.State {
	case StateStartScreen:
		if g.IsKeyJustPressed(ebiten.Key1){
			g.State = StateSinglePlayer
		}
		if g.IsKeyJustPressed(ebiten.Key2){
			g.State = StateMultiplayerMenu
		}
		
	case StateSinglePlayer :
		if g.IsKeyJustPressed(ebiten.KeyEscape){
			g.State = StateStartScreen 
			g.Score = 0
		}
		// edit multiplayer
		g.Paddle.MoveOnKeyPress()
		g.Ball.Move()
		g.CollideWithWall()
		g.CollideWithPaddle()

	case StateMultiplayerMenu:
		if g.IsKeyJustPressed(ebiten.KeyEscape){
			g.State = StateStartScreen 
		}
		if g.IsKeyJustPressed(ebiten.KeyH){
			// go network.StartServer() // goroutine doesnt block
			// g.State = StateConnecting
		}
		if g.IsKeyJustPressed(ebiten.KeyJ){
			// go network.ConnectToHost("127.0.0.1:8080") // hardcoded for now 
			// g.State = StateConnecting
		}
	case StateGameOver:
		// if g.IsKeyJustPressed(ebiten.Key1) {
		// 	g.State = StateSinglePlayer
		// }
		if g.IsKeyJustPressed(ebiten.Key2){
			g.State = StateStartScreen 
		}
		
	}
	return nil
}

func (g *Game) Layout(w, h int) (int, int) {
	return 640, 400
}

func (g *Game) IsKeyJustPressed(key ebiten.Key) bool {
	pressed := ebiten.IsKeyPressed(key)
	justPressed := pressed && !g.PrevKeys[key]
	g.PrevKeys[key] = pressed
	return justPressed
}

