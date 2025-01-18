package game

import (
	"image/color"

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
	greyColor = color.RGBA{82, 82, 82, 255}
)

func (g *Game) GenerateBoard(screen *ebiten.Image, sWidth int) *ebiten.Image {
	dc := gg.NewContext(sWidth, sWidth)
	dc.SetColor(color.Black)
	dc.Clear()

	// Draw grid lines
	dc.SetColor(color.White)
	for i := 1; i <= 2; i++ {
		gridLinePosition := i*cellSize - gridLineThickness/2
		dc.DrawLine(float64(gridLinePosition), 0, float64(gridLinePosition), float64(sWidth))
		dc.DrawLine(0, float64(gridLinePosition), float64(sWidth), float64(gridLinePosition))
	}
	dc.SetLineWidth(gridLineThickness)
	dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}

func (g *Game) GenerateSymbols(screen *ebiten.Image) (*ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image) {
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

	imageOHighlighted := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageOHighlighted.SetColor(greyColor)
	imageOHighlighted.DrawCircle(effectiveCellSize/2, effectiveCellSize/2, effectiveCellSize/2-symbolSpacing)
	imageOHighlighted.SetLineWidth(symbolThickness)
	imageOHighlighted.Stroke()

	imageXHighlighted := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageXHighlighted.SetColor(greyColor)
	imageXHighlighted.SetLineWidth(symbolThickness)
	imageXHighlighted.DrawLine(xLinesWidth, xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing)
	imageXHighlighted.DrawLine(xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing, xLinesWidth)
	imageXHighlighted.Stroke()

	imageEmpty := gg.NewContext(effectiveCellSize-3, effectiveCellSize-3)
	imageEmpty.SetColor(color.Black)
	imageEmpty.Clear()

	return ebiten.NewImageFromImage(imageX.Image()), ebiten.NewImageFromImage(imageO.Image()), ebiten.NewImageFromImage(imageXHighlighted.Image()), ebiten.NewImageFromImage(imageOHighlighted.Image()), ebiten.NewImageFromImage(imageEmpty.Image())
}
