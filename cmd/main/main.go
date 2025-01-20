// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// GoRythm is a Tic Tac Toe mini-game where 2 players have to place their symbols on the beat
// of the music. The players can win by aligning 3 symbols or having the most beat score at the end.
package main

import (
	_ "image/png"

	"GoRythm/game"
	a "GoRythm/internal/audio"
	"GoRythm/internal/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sWidth  = 480       // Screen width
	sHeight = 700       // Screen height
	title   = "GoRythm" // Window title
)

func main() {
	audioContext := audio.NewContext(a.SampleRate) // Initialize the audio context once

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
