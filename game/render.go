package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {

	if g.state == StateMenu {
		g.DrawMenu(screen)
		return
	}
	if g.state == StateLoading {
		g.DrawTimer(screen)
	}
	// Draw the game elements when playing
	if g.state == StatePlaying {
		g.DrawGame(screen)
	}
	if g.state == StateGameOver {
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

	switch g.gameMode {
	case 1:
		colorEasy = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	case 2:
		colorMedium = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	case 3:
		colorHard = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	}

	text.Draw(screen, "1. Easy", normalText, 70, 250, colorEasy)
	text.Draw(screen, "2. Hard", normalText, 70, 300, colorMedium)
	text.Draw(screen, "3. GoRythm", normalText, 70, 350, colorHard)

	msgStart := "Press ENTER to start"
	text.Draw(screen, msgStart, normalText, g.sWidth/2, g.sHeight/2, color.White)
}

func (g *Game) DrawTimer(screen *ebiten.Image) {
	// Make a countdown timer of 3 seconds
	if g.countdown > 0 {
		msgTimer := fmt.Sprintf("%v", g.countdown)
		if g.countdown == 0 {
			msgTimer = "Go"
		}
		text.Draw(screen, msgTimer, bigText, (g.sWidth-text.BoundString(bigText, msgTimer).Dx())/2, g.sHeight/2, color.White)
	}
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	screen.DrawImage(g.boardImage, nil)
	screen.DrawImage(g.gameImage, nil)

	// Calculate the elapsed time
	elapsed := time.Since(g.startTime).Seconds()

	if g.state != StateGameOver {
		if g.gameMode == 3 {
			for _, beat := range g.goRythm.beatMap {
				if elapsed >= beat.Time && elapsed < beat.Time+0.1 { // Allow a small margin for matching
					g.circleColorChangeTime = time.Now()
					break
				}
			}
		}

		// Draw the circle
		circleColor := color.RGBA{0, 0, 255, 255} // Blue color
		if time.Since(g.circleColorChangeTime).Seconds() < 0.5 {
			circleColor = color.RGBA{255, 0, 0, 255} // Red color
		}
		ebitenutil.DrawCircle(screen, float64(g.sWidth)/2, float64(g.sHeight)-100, 50, circleColor)
	}

	// Draw rounds
	msgRounds := fmt.Sprintf("Round: %v", g.rounds)
	text.Draw(screen, msgRounds, normalText, 10, g.sHeight-30, color.White)

	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	text.Draw(screen, msgOX, normalText, (g.sWidth-150)/2, g.sHeight-5, color.White)

	msgPlayer := fmt.Sprintf("Player: %v", g.playing)
	text.Draw(screen, msgPlayer, normalText, 10, g.sHeight-50, color.White)
}

func (g *Game) DrawGameOver(screen *ebiten.Image) {
	g.DrawGame(screen)
	if g.win != "" {
		_, winningLine := g.CheckWin()
		if winningLine != nil {
			dc := gg.NewContext(g.sWidth, g.sWidth)
			dc.SetColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}) // Red color for the winning line
			dc.SetLineWidth(10)
			startX := float64(winningLine[0][1]*160 + 80)
			startY := float64(winningLine[0][0]*160 + 80)
			endX := float64(winningLine[2][1]*160 + 80)
			endY := float64(winningLine[2][0]*160 + 80)
			dc.DrawLine(startX, startY, endX, endY)
			dc.Stroke()
			screen.DrawImage(ebiten.NewImageFromImage(dc.Image()), nil)
		}
	}
	msgPressEnter := "Press ENTER to play again"
	text.Draw(screen, msgPressEnter, normalText, (g.sWidth-150)/2, g.sHeight-30, color.White)
	if g.win != "" {
		msgWin := fmt.Sprintf("%v wins!", g.win)
		text.Draw(screen, msgWin, bigText, (g.sWidth-150)/2, g.sHeight-60, color.RGBA{G: 50, B: 200, A: 255})
	} else {
		msgDraw := "It's a draw!"
		text.Draw(screen, msgDraw, bigText, (g.sWidth-150)/2, g.sHeight-60, color.RGBA{G: 50, B: 200, A: 255})
	}
	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	text.Draw(screen, msgOX, normalText, (g.sWidth-150)/2, g.sHeight-5, color.White)
}
