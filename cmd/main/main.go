package main

import (
	_ "image/png"

	"GoRythm/game"
	"GoRythm/internal/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sWidth     = 480
	sHeight    = 700
	title      = "GoRythm"
	sampleRate = 44100
)

func main() {
	audioContext := audio.NewContext(sampleRate) // Initialize the audio context once

	// Initialize the game
	game := game.NewGame()
	err := game.Init(audioContext, sWidth, sHeight)
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to initialize the game: "+err.Error())
	}
	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle(title)

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.LogMessage(log.FATAL, "Failed to run the game: "+err.Error())
	}
}
