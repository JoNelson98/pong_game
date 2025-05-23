package main

import (
	"log"

	"github.com/JoNelson98/go_pong/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Pong in Ebitengine")
	ebiten.SetWindowSize(640, 400)

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
