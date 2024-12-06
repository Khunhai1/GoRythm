package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var boardImage *ebiten.Image
var gameImage = ebiten.NewImage(sWidth, sWidth)

func (g *Game) Init() {
	boardImage = generateBoardImage()
	re := newRandom().Intn(2)
	if re == 0 {
		g.playing = "O"
		g.alter = 0
	} else {
		g.playing = "X"
		g.alter = 1
	}
	g.Load()
	g.ResetPoints()
}

func (g *Game) Load() {
	gameImage.Clear()
	g.gameBoard = [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	g.round = 0
	if g.alter == 0 {
		g.playing = "X"
		g.alter = 1
	} else if g.alter == 1 {
		g.playing = "O"
		g.alter = 0
	}
	g.win = ""
	g.state = 1
}

func generateBoardImage() *ebiten.Image {
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
