package main

import (
	"math"
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
		g.board[x][y] = "O"
		opSymbol := &ebiten.DrawImageOptions{}
		opSymbol.GeoM.Translate(float64(x*160), float64(y*160))
		gameImage.DrawImage(OImage, opSymbol)
	case "X":
		g.board[x][y] = "X"
		opSymbol := &ebiten.DrawImageOptions{}
		opSymbol.GeoM.Translate(float64(x*160), float64(y*160))
		gameImage.DrawImage(XImage, opSymbol)
	}
	g.switchPlayer()
	g.rounds++
}

func (g *Game) switchPlayer() {
	if g.playing == "X" {
		g.playing = "O"
	} else {
		g.playing = "X"
	}
	if g.player == "human" {
		g.player = "ai"
	} else {
		g.player = "human"
	}
}

func (g *Game) EasyCpu() (int, int) {
	r := newRandom()
	var x, y int
	for {
		x = r.Intn(3)
		y = r.Intn(3)
		if g.board[x][y] == "" {
			break
		}
	}
	return x, y
}

// HardCpu playing with minimax algorithm
func (g *Game) HardCpu() (int, int) {
	bestScore := math.MinInt
	var bestMove [2]int
	// Minimax with increased depth
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if g.board[x][y] == "" {
				// Simulate the move
				g.board[x][y] = "X"
				score := g.minimax(0, false)
				g.board[x][y] = "" // Undo the move

				// Keep track of the best move
				if score > bestScore {
					bestScore = score
					bestMove = [2]int{x, y}
				}
			}
		}
	}
	return bestMove[0], bestMove[1]
}

func (g *Game) minimax(depth int, isMaximizing bool) int {
	// Check if game is over
	winner := g.CheckWin()
	if winner == "X" {
		return 10 - depth // Maximize for AI (X)
	}
	if winner == "O" {
		return depth - 10 // Minimize for Player (O)
	}
	if g.IsBoardFull() {
		return 0 // Draw
	}

	// Maximizing Player (AI)
	if isMaximizing {
		bestScore := math.MinInt
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if g.board[x][y] == "" {
					g.board[x][y] = "X" // AI's move
					score := g.minimax(depth+1, false)
					g.board[x][y] = "" // Undo the move
					bestScore = max(bestScore, score)
				}
			}
		}
		return bestScore
	} else { // Minimizing Player (Human)
		bestScore := math.MaxInt
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if g.board[x][y] == "" {
					g.board[x][y] = "O" // Human's move
					score := g.minimax(depth+1, true)
					g.board[x][y] = "" // Undo the move
					bestScore = min(bestScore, score)
				}
			}
		}
		return bestScore
	}
}

func (g *Game) CheckWin() string {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] != "" {
			return g.board[i][0]
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] != "" {
			return g.board[0][i]
		}
	}
	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] != "" {
		return g.board[0][0]
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] != "" {
		return g.board[0][2]
	}
	return ""
}

func (g *Game) IsBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == "" {
				return false
			}
		}
	}
	return true
}
