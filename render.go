package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {

	if gameState == StateMenu {
		g.DrawMenu(screen)
		return
	}
	if gameState == StateLoading {
		g.DrawTimer(screen)
	}
	// Draw the game elements when playing
	if gameState == StatePlaying {
		g.DrawGame(screen)
	}
	if gameState == StateGameOver {
		g.DrawGameOver(screen)
	}
}

func (g *Game) DrawMenu(screen *ebiten.Image) {
	msgTitle := "GopherTicTacToe"
	text.Draw(screen, msgTitle, bigText, 30, 100, color.White)
	msgDifficulty := "Choose difficulty:"
	text.Draw(screen, msgDifficulty, normalText, 70, 200, color.White)

	// Highlight the selected difficulty
	colorEasy := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	colorMedium := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	colorHard := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	switch g.difficulty {
	case 1:
		colorEasy = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	case 2:
		colorMedium = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	case 3:
		colorHard = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	}

	text.Draw(screen, "1. Easy", normalText, 70, 250, colorEasy)
	text.Draw(screen, "2. Medium", normalText, 70, 300, colorMedium)
	text.Draw(screen, "3. Hard", normalText, 70, 350, colorHard)

	msgStart := "Press ENTER to start"
	text.Draw(screen, msgStart, normalText, sWidth/2, sHeight/2, color.White)
}

func (g *Game) DrawTimer(screen *ebiten.Image) {
	// Make a countdown timer of 3 seconds
	if g.countdown > 0 {
		msgTimer := fmt.Sprintf("%v", g.countdown)
		if g.countdown == 0 {
			msgTimer = "Go"
		}
		text.Draw(screen, msgTimer, bigText, (sWidth-text.BoundString(bigText, msgTimer).Dx())/2, sHeight/2, color.White)
	}
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	screen.DrawImage(boardImage, nil)
	screen.DrawImage(gameImage, nil)
	// mx, my := ebiten.CursorPosition()

	msgFPS := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	text.Draw(screen, msgFPS, normalText, sWidth-100, 30, color.White)

	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	text.Draw(screen, msgOX, normalText, (sWidth-150)/2, sHeight-5, color.White)
	if g.win != "" {
		msgWin := fmt.Sprintf("%v wins!", g.win)
		text.Draw(screen, msgWin, bigText, 70, 200, color.RGBA{G: 50, B: 200, A: 255})
	}
}

func (g *Game) DrawGameOver(screen *ebiten.Image) {
	gameImage.Clear()
	boardImage.Clear()
	msgGameOver := "Game Over"
	text.Draw(screen, msgGameOver, bigText, 70, 200, color.White)
	msgPressEnter := "Press ENTER to play again"
	text.Draw(screen, msgPressEnter, normalText, 70, 300, color.White)
}
