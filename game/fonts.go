package game

import (
	"GoTicTacToe/internal/log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	fontSize    = 15
	bigFontSize = 50
	dpi         = 72
)

var normalText font.Face
var bigText font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to parse font: "+err.Error())
	}
	normalText, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to parse font: "+err.Error())
	}
	bigText, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    bigFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to parse font: "+err.Error())
	}
}
