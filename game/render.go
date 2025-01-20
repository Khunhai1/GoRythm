// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"GoRythm/internal/log"
	t "GoRythm/internal/text"
	"GoRythm/internal/theme"
	"fmt"
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
	msgTitle := "GoRythm"
	t.DrawText(screen, msgTitle, t.BigText, 30, 100, theme.TextColor)
	msgDifficulty := "Choose difficulty:"
	t.DrawText(screen, msgDifficulty, t.NormalText, 70, 200, theme.TextColor)

	// Highlight the selected difficulty
	colorClassic := theme.TextColor
	colorEasy := theme.TextColor
	colorHard := theme.TextColor
	colorGoRythm := theme.TextColor

	switch g.gameMode {
	case CLASSIC_PVP_MODE:
		colorClassic = theme.SelectedTextColor
	case EASY_AI_MODE:
		colorEasy = theme.SelectedTextColor
	case HARD_AI_MODE:
		colorHard = theme.SelectedTextColor
	case GORYTHM_MODE:
		colorGoRythm = theme.SelectedTextColor
	}

	t.DrawText(screen, "1. PVP - Classic", t.NormalText, 70, 250, colorClassic)
	t.DrawText(screen, "2. Easy", t.NormalText, 70, 300, colorEasy)
	t.DrawText(screen, "3. Hard", t.NormalText, 70, 350, colorHard)
	t.DrawText(screen, "4. GoRythm", t.NormalText, 70, 400, colorGoRythm)

	msgStart := "Press ENTER to start"
	t.DrawText(screen, msgStart, t.NormalText, g.sWidth/2, g.sHeight/2, theme.TextColor)
}

func (g *Game) DrawTimer(screen *ebiten.Image) {
	// Make a countdown timer of 3 seconds
	if g.countdown > 0 {
		msgTimer := fmt.Sprintf("%v", g.countdown)
		if g.countdown == 0 {
			msgTimer = "Go"
		}
		textWidth, _ := text.Measure(msgTimer, t.BigText, 0)
		t.DrawText(screen, msgTimer, t.BigText, (g.sWidth-int(textWidth))/2, g.sHeight/2, theme.TextColor)
	}
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	if g.boardImage == nil || g.gameImage == nil {
		log.LogMessage(log.FATAL, "boardImage or gameImage is nil")
	}
	screen.DrawImage(g.boardImage, nil)
	screen.DrawImage(g.gameImage, nil)

	if g.gameMode == GORYTHM_MODE {
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
			beat := false
			circleColor := theme.CircleNoBeatColor
			if time.Since(g.goRythm.circleColorChangeTime).Seconds() < 0.5 {
				circleColor = theme.CircleBeatColor
				beat = true
			}
			vector.DrawFilledCircle(screen, float32(g.sWidth)/2, float32(g.sHeight)-100, 50, circleColor, false)
			if beat {
				msgBeat := "Click !"
				textWidth, _ := text.Measure(msgBeat, t.NormalText, 0)
				t.DrawText(screen, msgBeat, t.NormalText, (g.sWidth-int(textWidth))/2, g.sHeight-100, theme.TextColor)
			}
		}
	}

	// Draw rounds
	msgRounds := fmt.Sprintf("Round: %v", g.rounds)
	t.DrawText(screen, msgRounds, t.NormalText, 10, g.sHeight-30, theme.TextColor)

	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	t.DrawText(screen, msgOX, t.NormalText, (g.sWidth-150)/2, g.sHeight-30, theme.TextColor)

	msgPlayer := fmt.Sprintf("Player: %v", g.currentPlayerSymbol)
	t.DrawText(screen, msgPlayer, t.NormalText, 10, g.sHeight-60, theme.TextColor)
}

func (g *Game) DrawGameOver(screen *ebiten.Image) {
	g.DrawGame(screen)
	if g.win != NONE_PLAYING || g.gameMode == GORYTHM_MODE {
		_, winningLine := g.CheckWinBoard()
		if winningLine != nil {
			dc := gg.NewContext(g.sWidth, g.sWidth)
			dc.SetColor(theme.WinningLineColor)
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
	t.DrawText(screen, msgPressEnter, t.NormalText, (g.sWidth-150)/2, g.sHeight-130, theme.TextColor)
	if g.win != NONE_PLAYING {
		msgWin := fmt.Sprintf("%v wins!", g.win)
		t.DrawText(screen, msgWin, t.BigText, (g.sWidth-150)/2, g.sHeight-100, theme.GameOverTextColor)
	} else if g.gameMode == GORYTHM_MODE {
		msgDraw := "Score draw!"
		t.DrawText(screen, msgDraw, t.BigText, (g.sWidth-150)/2, g.sHeight-100, theme.GameOverTextColor)
	} else {
		msgDraw := "It's a draw!"
		t.DrawText(screen, msgDraw, t.BigText, (g.sWidth-150)/2, g.sHeight-100, theme.GameOverTextColor)
	}
	msgOX := fmt.Sprintf("O Score: %v | X Score: %v", g.pointsO, g.pointsX)
	t.DrawText(screen, msgOX, t.NormalText, (g.sWidth-150)/2, g.sHeight-30, theme.TextColor)
}
