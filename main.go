package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	sWidth  = 480
	sHeight = 600
)

func main() {
	game := &Game{}
	err := game.Init()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("GopherTicTacToe")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
