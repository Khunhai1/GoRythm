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
	// Get the current elapsed time
	elapsed := time.Since(g.startTime).Seconds()

	// Find the closest beat time
	var closestBeatTime float64
	minDifference := math.MaxFloat64
	for _, beat := range g.goRythm.beatMap {
		difference := math.Abs(beat.Time - elapsed)
		if difference < minDifference {
			minDifference = difference
			closestBeatTime = beat.Time
		}
	}

	// Calculate the precision score
	precisionScore := g.CalculateScore(closestBeatTime)

	switch g.playing {
	case "O":
		g.board[x][y] = "O"
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
		gameImage.DrawImage(OImage, options)
		g.pointsO += precisionScore
	case "X":
		g.board[x][y] = "X"
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
		gameImage.DrawImage(XImage, options)
		g.pointsX += precisionScore
	}
	g.switchPlayer()
	g.rounds++
}

func (g *Game) removeSymbol(x, y int) {
	g.board[x][y] = ""
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
	gameImage.DrawImage(EmptyImage, options)
}

func (g *Game) highlightSymbol(x, y int) {
	switch g.playing {
	case "O":
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
		gameImage.DrawImage(OImageHighlighted, options)
	case "X":
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
		gameImage.DrawImage(XImageHighlighted, options)
	}
}

func (g *Game) switchPlayer() {
	if g.playing == "X" {
		g.playing = "O"
	} else {
		g.playing = "X"
	}
	if g.gameMode != 3 {
		if g.player == "human" {
			g.player = "ai"
		} else {
			g.player = "human"
		}
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
	winner, _ := g.CheckWin()
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

func (g *Game) CheckWin() (string, [][]int) {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] != "" {
			return g.board[i][0], [][]int{{0, i}, {1, i}, {2, i}}
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] != "" {
			return g.board[0][i], [][]int{{i, 0}, {i, 1}, {i, 2}}
		}
	}
	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] != "" {
		return g.board[0][0], [][]int{{0, 0}, {1, 1}, {2, 2}}
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] != "" {
		return g.board[0][2], [][]int{{0, 2}, {1, 1}, {2, 0}}
	}
	return "", nil
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
