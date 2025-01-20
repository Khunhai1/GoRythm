// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"GoRythm/internal/theme"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellSize          = 160                            // Cell size in pixels (square)
	gridLineThickness = 2                              // Grid line thickness in pixels
	effectiveCellSize = cellSize - gridLineThickness/2 // Effective cell size in pixels (square) without grid line taking some space
	symbolThickness   = 15                             // Symbol thickness in pixels
	xLinesWidth       = 20                             // X symbol lines width in pixels
	symbolSpacing     = 20                             // Symbol spacing in pixels
)

func (g *Game) GenerateBoard(screen *ebiten.Image, sWidth int) *ebiten.Image {
	dc := gg.NewContext(sWidth, sWidth)
	dc.SetColor(theme.BackgroundColor)
	dc.Clear()

	// Draw grid lines
	dc.SetColor(theme.BoardColor)
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
	imageO.SetColor(theme.SymbolOColor)
	imageO.DrawCircle(effectiveCellSize/2, effectiveCellSize/2, effectiveCellSize/2-symbolSpacing)
	imageO.SetLineWidth(symbolThickness)
	imageO.Stroke()

	imageX := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageX.SetColor(theme.SymbolXColor)
	imageX.SetLineWidth(symbolThickness)
	imageX.DrawLine(xLinesWidth, xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing)
	imageX.DrawLine(xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing, xLinesWidth)
	imageX.Stroke()

	imageOHighlighted := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageOHighlighted.SetColor(theme.ToBeDeletedSymbolsColor)
	imageOHighlighted.DrawCircle(effectiveCellSize/2, effectiveCellSize/2, effectiveCellSize/2-symbolSpacing)
	imageOHighlighted.SetLineWidth(symbolThickness)
	imageOHighlighted.Stroke()

	imageXHighlighted := gg.NewContext(effectiveCellSize, effectiveCellSize)
	imageXHighlighted.SetColor(theme.ToBeDeletedSymbolsColor)
	imageXHighlighted.SetLineWidth(symbolThickness)
	imageXHighlighted.DrawLine(xLinesWidth, xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing)
	imageXHighlighted.DrawLine(xLinesWidth, effectiveCellSize-symbolSpacing, effectiveCellSize-symbolSpacing, xLinesWidth)
	imageXHighlighted.Stroke()

	imageEmpty := gg.NewContext(effectiveCellSize-3, effectiveCellSize-3)
	imageEmpty.SetColor(theme.BackgroundColor)
	imageEmpty.Clear()

	return ebiten.NewImageFromImage(imageX.Image()), ebiten.NewImageFromImage(imageO.Image()), ebiten.NewImageFromImage(imageXHighlighted.Image()), ebiten.NewImageFromImage(imageOHighlighted.Image()), ebiten.NewImageFromImage(imageEmpty.Image())
}
