// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"math/rand"
	"time"

	board "GoRythm/internal/generation"

	"github.com/hajimehoshi/ebiten/v2"
)

// newRandom returns a new random number generator based on the current time.
func newRandom() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}

// placeSymbol places the current player symbol on the board at the given position.
// It also calls the draw function to display the symbol on the screen.
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

// removeSymbol removes the symbol from the board at the given position.
// It also calls the draw function to remove the symbol from the screen.
func (g *Game) removeSymbol(x, y int) {
	g.board[x][y] = NONE_PLAYING
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x*board.CellSize), float64(y*board.CellSize))
	g.gameImage.DrawImage(g.EmptyImage, options)
}

// highlightSymbol highlights the symbol on the board at the given position.
// It also calls the draw function to display the highlighted symbol on the screen.
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

// switchPlayer switches the current player and player type based on the game mode.
func (g *Game) switchPlayer() {
	if g.currentPlayerSymbol == X_PLAYING {
		g.currentPlayerSymbol = O_PLAYING
	} else {
		g.currentPlayerSymbol = X_PLAYING
	}
	if g.gameMode != GORYTHM_MODE && g.gameMode != CLASSIC_PVP_MODE {
		if g.currentPlayerType == HUMAN_TYPE {
			g.currentPlayerType = AI_TYPE
		} else {
			g.currentPlayerType = HUMAN_TYPE
		}
	}
}

// checkWinBoard checks the winner based on the board and returns the winner and the winning positions.
func (g *Game) checkWinBoard() (winner SymbolPlaying, position [][]int) {
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

// checkWinScore checks the winner based on the score and returns the winner
func (g *Game) checkWinScore() (winner SymbolPlaying) {
	if g.pointsO > g.pointsX {
		return O_PLAYING
	} else if g.pointsX > g.pointsO {
		return X_PLAYING
	}
	return NONE_PLAYING
}

// isBoardFull checks if the board is full and returns true if it is.
func (g *Game) isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == NONE_PLAYING {
				return false
			}
		}
	}
	return true
}
