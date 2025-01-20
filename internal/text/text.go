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
	fontSize    = 15
	bigFontSize = 50
	dpi         = 72
)

var (
	NormalText text.Face
	BigText    text.Face
	font       []byte = fonts.MPlus1pRegular_ttf
)

func init() {
	fontSrc, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.LogMessage(log.FATAL, "Failed to parse font: "+err.Error())
	}
	NormalText = &text.GoTextFace{
		Source: fontSrc,
		Size:   fontSize,
	}
	BigText = &text.GoTextFace{
		Source: fontSrc,
		Size:   bigFontSize,
	}
}

func DrawText(screen *ebiten.Image, msg string, font text.Face, x, y int, color color.Color) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color)
	text.Draw(screen, msg, font, options)
}
