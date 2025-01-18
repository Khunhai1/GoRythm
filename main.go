package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sWidth  = 480
	sHeight = 700
)

var audioContext *audio.Context

func main() {
	audioContext = audio.NewContext(44100) // Initialize the audio context once

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
