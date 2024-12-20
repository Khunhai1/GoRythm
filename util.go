package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func newRandom() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}

func (g *Game) ResetPoints() {
	g.pointsO = 0
	g.pointsX = 0
}

func (g *Game) placeSymbol(x int, y int) {
	switch g.playing {
	case "O":
		if g.board[x/160][y/160] == "" {
			g.board[x/160][y/160] = "O"
			opSymbol := &ebiten.DrawImageOptions{}
			opSymbol.GeoM.Translate(float64(x), float64(y))
			gameImage.DrawImage(OImage, opSymbol)
			g.switchPlayer()
			g.player = "ai"
		}
	case "X":
		if g.board[x/160][y/160] == "" {
			g.board[x/160][y/160] = "X"
			opSymbol := &ebiten.DrawImageOptions{}
			opSymbol.GeoM.Translate(float64(x), float64(y))
			gameImage.DrawImage(XImage, opSymbol)
			g.switchPlayer()
			g.player = "ai"
		}
	}
	g.rounds++
}

func (g *Game) switchPlayer() {
	if g.playing == "O" {
		g.playing = "X"
		g.alter = 1
	} else {
		g.playing = "O"
		g.alter = 0
	}
}

// Check if the game is over
func (g *Game) CheckWin() {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] != "" {
			g.win = g.board[i][0]
			return
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] != "" {
			g.win = g.board[0][i]
			return
		}
	}
	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] != "" {
		g.win = g.board[0][0]
		return
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] != "" {
		g.win = g.board[0][2]
		return
	}
}

func (g *Game) EasyCpu() {
	r := newRandom()
	var x, y int
	for {
		x = r.Intn(3)
		y = r.Intn(3)
		if g.board[x][y] == "" {
			break
		}
	}
	g.board[x][y] = g.playing
	opSymbol := &ebiten.DrawImageOptions{}
	opSymbol.GeoM.Translate(float64(x*160), float64(y*160))
	if g.playing == "O" {
		gameImage.DrawImage(OImage, opSymbol)
	} else {
		gameImage.DrawImage(XImage, opSymbol)
	}
	g.switchPlayer()
}
