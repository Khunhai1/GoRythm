// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// Package generation contains the board elements generation functions for the Tic-Tac-Toe game
// as well as the constants and variables used to draw those elements.
package generation

import (
	"GoRythm/internal/theme"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CellSize          = 160                            // Cell size in pixels (square)
	GridLineThickness = 2                              // Grid line thickness in pixels
	EffectiveCellSize = CellSize - GridLineThickness/2 // Effective cell size in pixels (square) without grid line taking some space
	SymbolThickness   = 15                             // Symbol thickness in pixels
	XLinesWidth       = 20                             // X symbol lines width in pixels
	SymbolSpacing     = 20                             // Symbol spacing in pixels
	EmptyImageSize    = EffectiveCellSize - 3          // Empty image size in pixels that is used to remove symbols in GoRythm mode (square a bit smaller than the cell)
)

// GenerateBoard generates the board image with the grid lines and returns it.
func GenerateBoard(screen *ebiten.Image, sWidth int) *ebiten.Image {
	dc := gg.NewContext(sWidth, sWidth)
	dc.SetColor(theme.BackgroundColor)
	dc.Clear()

	// Draw grid lines
	dc.SetColor(theme.BoardColor)
	for i := 1; i <= 2; i++ {
		gridLinePosition := i*CellSize - GridLineThickness/2
		dc.DrawLine(float64(gridLinePosition), 0, float64(gridLinePosition), float64(sWidth))
		dc.DrawLine(0, float64(gridLinePosition), float64(sWidth), float64(gridLinePosition))
	}
	dc.SetLineWidth(GridLineThickness)
	dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}

// GenerateSymbols generates the symbols images (X, O, highlighted X, highlighted O and empty) and returns them.
func GenerateSymbols(screen *ebiten.Image) (*ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image) {
	dc := gg.NewContext(EffectiveCellSize, EffectiveCellSize)
	dc.Clear()

	// O symbol image
	imageO := gg.NewContext(EffectiveCellSize, EffectiveCellSize)
	imageO.SetColor(theme.SymbolOColor)
	imageO.DrawCircle(EffectiveCellSize/2, EffectiveCellSize/2, EffectiveCellSize/2-SymbolSpacing)
	imageO.SetLineWidth(SymbolThickness)
	imageO.Stroke()

	// X symbol image
	imageX := gg.NewContext(EffectiveCellSize, EffectiveCellSize)
	imageX.SetColor(theme.SymbolXColor)
	imageX.SetLineWidth(SymbolThickness)
	imageX.DrawLine(XLinesWidth, XLinesWidth, EffectiveCellSize-SymbolSpacing, EffectiveCellSize-SymbolSpacing)
	imageX.DrawLine(XLinesWidth, EffectiveCellSize-SymbolSpacing, EffectiveCellSize-SymbolSpacing, XLinesWidth)
	imageX.Stroke()

	// Highlighted O symbol image
	imageOHighlighted := gg.NewContext(EffectiveCellSize, EffectiveCellSize)
	imageOHighlighted.SetColor(theme.ToBeDeletedSymbolsColor)
	imageOHighlighted.DrawCircle(EffectiveCellSize/2, EffectiveCellSize/2, EffectiveCellSize/2-SymbolSpacing)
	imageOHighlighted.SetLineWidth(SymbolThickness)
	imageOHighlighted.Stroke()

	// Highlighted X symbol image
	imageXHighlighted := gg.NewContext(EffectiveCellSize, EffectiveCellSize)
	imageXHighlighted.SetColor(theme.ToBeDeletedSymbolsColor)
	imageXHighlighted.SetLineWidth(SymbolThickness)
	imageXHighlighted.DrawLine(XLinesWidth, XLinesWidth, EffectiveCellSize-SymbolSpacing, EffectiveCellSize-SymbolSpacing)
	imageXHighlighted.DrawLine(XLinesWidth, EffectiveCellSize-SymbolSpacing, EffectiveCellSize-SymbolSpacing, XLinesWidth)
	imageXHighlighted.Stroke()

	// Empty symbol image
	imageEmpty := gg.NewContext(EmptyImageSize, EmptyImageSize)
	imageEmpty.SetColor(theme.BackgroundColor)
	imageEmpty.Clear()

	return ebiten.NewImageFromImage(imageX.Image()), ebiten.NewImageFromImage(imageO.Image()), ebiten.NewImageFromImage(imageXHighlighted.Image()), ebiten.NewImageFromImage(imageOHighlighted.Image()), ebiten.NewImageFromImage(imageEmpty.Image())
}
