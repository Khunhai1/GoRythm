package main

import (
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var gameImage = ebiten.NewImage(sWidth, sWidth)
var XImage, OImage *ebiten.Image
var boardImage *ebiten.Image

func (g *Game) Init() error {
	// Reset variables before starting a new game
	g.board = [3][3]string{}     // Reset the game board
	g.rounds = 0                 // Reset the number of rounds
	g.win = ""                   // Reset the win status
	g.playing = ""               // Reset the current player
	g.gameMode = 0               // Reset the game mode
	g.countdown = 3              // Reset the countdown timer
	g.countdownTime = time.Now() // Reset the countdown start time

	// Generate the game board and symbols
	boardImage = g.GenerateBoard(gameImage)
	XImage, OImage = g.GenerateSymbols(gameImage)
	re := newRandom().Intn(2)
	if re == 0 {
		g.playing = "O"
		g.alter = 0
	} else {
		g.playing = "X"
		g.alter = 1
	}

	// Initialize audio settings
	err := g.initAudio()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) GenerateBoard(screen *ebiten.Image) *ebiten.Image {
	const gridSize = 160
	dc := gg.NewContext(sWidth, sWidth)
	dc.SetColor(color.Black)
	dc.Clear()

	// Draw grid lines
	dc.SetColor(color.White)
	for i := 1; i <= 2; i++ {
		dc.DrawLine(float64(i*gridSize), 0, float64(i*gridSize), sWidth)
		dc.DrawLine(0, float64(i*gridSize), sWidth, float64(i*gridSize))
	}
	dc.SetLineWidth(5)
	dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}

func (g *Game) GenerateSymbols(screen *ebiten.Image) (*ebiten.Image, *ebiten.Image) {
	const gridSize = 160
	dc := gg.NewContext(gridSize, gridSize)
	dc.Clear()

	imageO := gg.NewContext(gridSize, gridSize)
	imageO.SetColor(color.White)
	imageO.DrawCircle(gridSize/2, gridSize/2, gridSize/2-10)
	imageO.SetLineWidth(15)
	imageO.Stroke()

	imageX := gg.NewContext(gridSize, gridSize)
	imageX.SetColor(color.White)
	imageX.SetLineWidth(15)
	imageX.DrawLine(20, 20, gridSize-20, gridSize-20)
	imageX.DrawLine(20, gridSize-20, gridSize-20, 20)
	imageX.Stroke()

	return ebiten.NewImageFromImage(imageX.Image()), ebiten.NewImageFromImage(imageO.Image())
}

// Init the audio player
func (g *Game) initAudio() error {
	if g.audioPlayer != nil {
		g.audioPlayer.Stop()
	}
	ap, err := NewAudioPlayer(audioContext) // Use the global audio context
	if err != nil {
		log.Printf("failed to init audio player: %v", err)
		return err
	}
	g.audioPlayer = ap
	return nil
}
