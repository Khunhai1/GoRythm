package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func keyChangeColor(key ebiten.Key, screen *ebiten.Image) {
	if inpututil.KeyPressDuration(key) > 1 {
		var msgText string
		var colorText color.RGBA
		colorChange := 255 - (255 / 60 * uint8(inpututil.KeyPressDuration(key)))
		if key == ebiten.KeyEscape {
			msgText = fmt.Sprintf("CLOSING...")
			colorText = color.RGBA{R: 255, G: colorChange, B: colorChange, A: 255}
		} else if key == ebiten.KeyR {
			msgText = fmt.Sprintf("RESETING...")
			colorText = color.RGBA{R: colorChange, G: 255, B: 255, A: 255}
		}
		text.Draw(screen, msgText, normalText, sWidth/2, sHeight-30, colorText)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(boardImage, nil)
	screen.DrawImage(gameImage, nil)
	mx, my := ebiten.CursorPosition()

	msgFPS := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	text.Draw(screen, msgFPS, normalText, 0, sHeight-30, color.White)

	keyChangeColor(ebiten.KeyEscape, screen)
	keyChangeColor(ebiten.KeyR, screen)
	msgOX := fmt.Sprintf("O: %v | X: %v", g.pointsO, g.pointsX)
	text.Draw(screen, msgOX, normalText, sWidth/2, sHeight-5, color.White)
	if g.win != "" {
		msgWin := fmt.Sprintf("%v wins!", g.win)
		text.Draw(screen, msgWin, bigText, 70, 200, color.RGBA{G: 50, B: 200, A: 255})
	}
	msg := fmt.Sprintf("%v", g.playing)
	text.Draw(screen, msg, normalText, mx, my, color.RGBA{G: 255, A: 255})
}

func (g *Game) DrawSymbol(x, y int, sym string) {
	const gridSize = 160
	dc := gg.NewContext(gridSize, gridSize)
	dc.Clear()

	// Draw O or X
	if sym == "O" {
		dc.SetColor(color.White)
		dc.DrawCircle(gridSize/2, gridSize/2, gridSize/2-10)
		dc.SetLineWidth(15)
		dc.Stroke()
	} else if sym == "X" {
		dc.SetColor(color.White)
		dc.SetLineWidth(15)
		dc.DrawLine(20, 20, gridSize-20, gridSize-20)
		dc.DrawLine(20, gridSize-20, gridSize-20, 20)
		dc.Stroke()
	}

	// Translate the symbol to the appropriate grid position
	opSymbol := &ebiten.DrawImageOptions{}
	opSymbol.GeoM.Translate(float64(x*gridSize), float64(y*gridSize))
	gameImage.DrawImage(ebiten.NewImageFromImage(dc.Image()), opSymbol)
}
