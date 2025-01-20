package game

import (
	"GoRythm/internal/log"
	t "GoRythm/internal/text"
	"fmt"
	"image/color"
	"time"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	t.DrawText(screen, msgTitle, t.BigText, 30, 100, color.White)
	msgDifficulty := "Choose difficulty:"
	t.DrawText(screen, msgDifficulty, t.NormalText, 70, 200, color.White)

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

	t.DrawText(screen, "1. Easy", t.NormalText, 70, 250, colorEasy)
	t.DrawText(screen, "2. Hard", t.NormalText, 70, 300, colorMedium)
	t.DrawText(screen, "3. GoRythm", t.NormalText, 70, 350, colorHard)

	msgStart := "Press ENTER to start"
	t.DrawText(screen, msgStart, t.NormalText, g.sWidth/2, g.sHeight/2, color.White)
}

func (g *Game) DrawTimer(screen *ebiten.Image) {
	// Make a countdown timer of 3 seconds
	if g.countdown > 0 {
		msgTimer := fmt.Sprintf("%v", g.countdown)
		if g.countdown == 0 {
			msgTimer = "Go"
		}
		textWidth, _ := text.Measure(msgTimer, t.BigText, 0)
		t.DrawText(screen, msgTimer, t.BigText, (g.sWidth-int(textWidth))/2, g.sHeight/2, color.White)
	}
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	if g.boardImage == nil || g.gameImage == nil {
		log.LogMessage(log.FATAL, "boardImage or gameImage is nil")
	}
	screen.DrawImage(g.boardImage, nil)
	screen.DrawImage(g.gameImage, nil)

	if g.gameMode == 3 {
		// Calculate the elapsed time
		elapsed := time.Since(g.goRythm.startTime).Seconds()

		if g.state != StateGameOver {
			for _, beat := range g.goRythm.beatMap {
				if elapsed >= beat.Time && elapsed < beat.Time+0.1 { // Allow a small margin for matching
					g.goRythm.circleColorChangeTime = time.Now()
					break
				}
			}

			// Draw the circle
			circleColor := color.RGBA{0, 0, 255, 255} // Blue color
			if time.Since(g.goRythm.circleColorChangeTime).Seconds() < 0.5 {
				circleColor = color.RGBA{255, 0, 0, 255} // Red color
			}
			vector.DrawFilledCircle(screen, float32(g.sWidth)/2, float32(g.sHeight)-100, 50, circleColor, false)
		}
	}

	// Draw rounds
	msgRounds := fmt.Sprintf("Round: %v", g.rounds)
	t.DrawText(screen, msgRounds, t.NormalText, 10, g.sHeight-30, color.White)

	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	t.DrawText(screen, msgOX, t.NormalText, (g.sWidth-150)/2, g.sHeight-5, color.White)

	msgPlayer := fmt.Sprintf("Player: %v", g.playing)
	t.DrawText(screen, msgPlayer, t.NormalText, 10, g.sHeight-50, color.White)
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
	t.DrawText(screen, msgPressEnter, t.NormalText, (g.sWidth-150)/2, g.sHeight-30, color.White)
	if g.win != "" {
		msgWin := fmt.Sprintf("%v wins!", g.win)
		t.DrawText(screen, msgWin, t.BigText, (g.sWidth-150)/2, g.sHeight-100, color.RGBA{G: 50, B: 200, A: 255})
	} else {
		msgDraw := "It's a draw!"
		t.DrawText(screen, msgDraw, t.BigText, (g.sWidth-150)/2, g.sHeight-100, color.RGBA{G: 50, B: 200, A: 255})
	}
	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	t.DrawText(screen, msgOX, t.NormalText, (g.sWidth-150)/2, g.sHeight-5, color.White)
}
