// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import "math"

// EasyCpu returns a random move.
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

// HardCpu returns the best move for the AI by using the minimax algorithm.
// It simulates all possible moves and passes them to the minimax function to find the best one.
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

// minimax is a recursive algorithm to find the best move possible.
// It returns the best score for the current player and board state.
func (g *Game) minimax(depth int, isMaximizing bool) int {
	// Check if game is over
	winner, _ := g.checkWinBoard()
	if winner == X_PLAYING {
		return 10 - depth // Maximize for AI (X)
	}
	if winner == O_PLAYING {
		return depth - 10 // Minimize for Player (O)
	}
	if g.isBoardFull() {
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
