// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// Package text provides the utiliies to draw text on the screen as well
// as the fonts used in the application.
package text

import (
	"GoRythm/internal/log"
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	normalFontSize = 15
	bigFontSize    = 50
)

var (
	NormalText text.Face
	BigText    text.Face
	font       []byte = fonts.MPlus1pRegular_ttf
)

// Initializes the fonts at the the initialization of the package.
func init() {
	fontSrc, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to parse font: "+err.Error())
	}
	NormalText = &text.GoTextFace{
		Source: fontSrc,
		Size:   normalFontSize,
	}
	BigText = &text.GoTextFace{
		Source: fontSrc,
		Size:   bigFontSize,
	}
}

// Draws text on the screen at the specified position, with the specified face and color.
func DrawText(screen *ebiten.Image, msg string, font text.Face, x, y int, color color.Color) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	text.Draw(screen, msg, font, options)
}
