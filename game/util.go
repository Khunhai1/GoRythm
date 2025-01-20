// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"math"
	"math/rand"
	"time"

	board "GoRythm/internal/generation"

	"github.com/hajimehoshi/ebiten/v2"
)

func newRandom() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}

func (g *Game) placeSymbol(x int, y int) {
	switch g.currentPlayerSymbol {
	case O_PLAYING:
		g.board[x][y] = O_PLAYING
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
		g.gameImage.DrawImage(g.OImage, options)
	case X_PLAYING:
		g.board[x][y] = X_PLAYING
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
		g.gameImage.DrawImage(g.XImage, options)
	}
}

func (g *Game) removeSymbol(x, y int) {
	g.board[x][y] = NONE_PLAYING
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
	g.gameImage.DrawImage(g.EmptyImage, options)
}

func (g *Game) highlightSymbol(x, y int) {
	switch g.currentPlayerSymbol {
	case O_PLAYING:
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
		g.gameImage.DrawImage(g.OImageHighlighted, options)
	case X_PLAYING:
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
		g.gameImage.DrawImage(g.XImageHighlighted, options)
	}
}

func (g *Game) switchPlayer() {
	if g.currentPlayerSymbol == X_PLAYING {
		g.currentPlayerSymbol = O_PLAYING
	} else {
		g.currentPlayerSymbol = X_PLAYING
	}
	if g.gameMode != GORYTHM_MODE {
		if g.currentPlayerType == HUMAN_TYPE {
			g.currentPlayerType = AI_TYPE
		} else {
			g.currentPlayerType = HUMAN_TYPE
		}
	}
}

func (g *Game) EasyCpu() (int, int) {
	r := newRandom()
	var x, y int
	for {
		x = r.Intn(3)
		y = r.Intn(3)
		if g.board[x][y] == NONE_PLAYING {
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
			if g.board[x][y] == NONE_PLAYING {
				// Simulate the move
				g.board[x][y] = X_PLAYING
				score := g.minimax(0, false)
				g.board[x][y] = NONE_PLAYING // Undo the move

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
	winner, _ := g.CheckWinBoard()
	if winner == X_PLAYING {
		return 10 - depth // Maximize for AI (X)
	}
	if winner == O_PLAYING {
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
				if g.board[x][y] == NONE_PLAYING {
					g.board[x][y] = X_PLAYING // AI's move
					score := g.minimax(depth+1, false)
					g.board[x][y] = NONE_PLAYING // Undo the move
					bestScore = max(bestScore, score)
				}
			}
		}
		return bestScore
	} else { // Minimizing Player (Human)
		bestScore := math.MaxInt
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if g.board[x][y] == NONE_PLAYING {
					g.board[x][y] = O_PLAYING // Human's move
					score := g.minimax(depth+1, true)
					g.board[x][y] = NONE_PLAYING // Undo the move
					bestScore = min(bestScore, score)
				}
			}
		}
		return bestScore
	}
}

func (g *Game) CheckWinBoard() (winner SymbolPlaying, position [][]int) {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] != NONE_PLAYING {
			return g.board[i][0], [][]int{{0, i}, {1, i}, {2, i}}
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] != NONE_PLAYING {
			return g.board[0][i], [][]int{{i, 0}, {i, 1}, {i, 2}}
		}
	}
	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] != NONE_PLAYING {
		return g.board[0][0], [][]int{{0, 0}, {1, 1}, {2, 2}}
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] != NONE_PLAYING {
		return g.board[0][2], [][]int{{0, 2}, {1, 1}, {2, 0}}
	}
	return NONE_PLAYING, nil
}

// CheckWinScore checks the winner based on the score and returns the winner
func (g *Game) CheckWinScore() (winner SymbolPlaying) {
	if g.pointsO > g.pointsX {
		return O_PLAYING
	} else if g.pointsX > g.pointsO {
		return X_PLAYING
	}
	return NONE_PLAYING
}

func (g *Game) IsBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == NONE_PLAYING {
				return false
			}
		}
	}
	return true
}
