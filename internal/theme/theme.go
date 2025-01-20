// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// Package theme provides the colors used in the game.
package theme

import "image/color"

var (
	BackgroundColor         color.Color = color.Black                            // Black
	TextColor               color.Color = color.White                            // White
	SelectedTextColor       color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	BoardColor              color.Color = color.White                            // White
	CircleNoBeatColor       color.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // Blue
	CircleBeatColor         color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	WinningLineColor        color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	GameOverTextColor       color.Color = color.RGBA{G: 50, B: 200, A: 255}      // Dark blue
	ToBeDeletedSymbolsColor color.Color = color.RGBA{82, 82, 82, 255}            // Grey
	SymbolXColor            color.Color = color.White                            // White
	SymbolOColor            color.Color = color.White                            // White
)
