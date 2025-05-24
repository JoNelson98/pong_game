package main

import (
	"log"

	"github.com/JoNelson98/go_pong/internal/config"
	"github.com/JoNelson98/go_pong/internal/game"
	"github.com/JoNelson98/go_pong/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	sound.Init()
	ebiten.SetWindowTitle("Pong in Ebitengine")
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
