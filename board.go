package main

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellSize          = 160
	gridLineThickness = 2
	effectiveCellSize = cellSize - gridLineThickness/2
	symbolThickness   = 15
	xLinesWidth       = 20
	symbolSpacing     = 20
)

var (
	gameImage, boardImage, XImage, OImage, EmptyImage *ebiten.Image
)

func (g *Game) Init() error {
	// Reset variables before starting a new game
	g.board = [3][3]string{} // Reset the game board
	g.rounds = 0             // Reset the number of rounds
	g.win = ""               // Reset the win status
	g.playing = ""           // Reset the current player
	g.gameMode = 0           // Reset the game mode
	g.countdown = 3          // Reset the countdown timer
	g.pointsO = 0            // Reset the points for O
	g.pointsX = 0            // Reset the points for X

	// Generate the game board and symbols
	gameImage = ebiten.NewImage(sWidth, sWidth)
	boardImage = g.GenerateBoard(gameImage)
	XImage, OImage, EmptyImage = g.GenerateSymbols(gameImage)
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
	loadBeatmap()
	return nil
}

func (g *Game) GenerateBoard(screen *ebiten.Image) *ebiten.Image {
	dc := gg.NewContext(sWidth, sWidth)
	dc.SetColor(color.Black)
	dc.Clear()

	// Draw grid lines
	dc.SetColor(color.White)
	for i := 1; i <= 2; i++ {
		gridLinePosition := i*cellSize - gridLineThickness/2
		dc.DrawLine(float64(gridLinePosition), 0, float64(gridLinePosition), sWidth)
		dc.DrawLine(0, float64(gridLinePosition), sWidth, float64(gridLinePosition))
	}
	dc.SetLineWidth(gridLineThickness)
	dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}

func (g *Game) GenerateSymbols(screen *ebiten.Image) (*ebiten.Image, *ebiten.Image, *ebiten.Image) {
	dc := gg.NewContext(effectiveCellSize, effectiveCellSize)
	dc.Clear()

	imageO := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageO.SetColor(color.White)
	imageO.DrawCircle(effectiveCellSize/2, effectiveCellSize/2, effectiveCellSize/2-symbolSpacing)
	imageO.SetLineWidth(symbolThickness)
	imageO.Stroke()

	imageX := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageX.SetColor(color.White)
	imageX.SetLineWidth(symbolThickness)
	imageX.DrawLine(xLinesWidth, xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing)
	imageX.DrawLine(xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing, xLinesWidth)
	imageX.Stroke()

	imageEmpty := gg.NewContext(effectiveCellSize-3, effectiveCellSize-3)
	imageEmpty.SetColor(color.Black)
	imageEmpty.Clear()

	return ebiten.NewImageFromImage(imageX.Image()), ebiten.NewImageFromImage(imageO.Image()), ebiten.NewImageFromImage(imageEmpty.Image())
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
