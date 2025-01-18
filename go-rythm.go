package main

import (
	"math"
	"time"
)

const (
	missedScore  = 0    // Score when missing a beat
	perfectScore = 300  // Score when hitting a beat perfectly
	goodScore    = 100  // Score when hitting a beat well
	okScore      = 50   // Score when hitting a beat ok
	perfectPrec  = 0.1  // Precision for perfect score (in seconds)
	goodPrec     = 0.25 // Precision for good score (in seconds)
	okPrec       = 0.4  // Precision for ok score (in seconds)
)

var (
	noMove = [2]int{-1, -1}
)

// A maximum of three symbols per player can be placed on the board. When the third symbol is placed, the first symbol is removed.
// The next symbol to be removed in the next round is highlighted.
type GoRythm struct {
	movesO       chan [2]int // The last two moves made by player O
	movesX       chan [2]int // The last two moves made by player X
	toBeRemovedO [2]int      // The last third move made by player O that will be removed next round
	toBeRemovedX [2]int      // The last third move made by player X that will be removed next round
	beatMap      []Beat      // The beat map for the music, containing the time and beat number of each beat
}

func NewGoRythmGame() *GoRythm {
	bm, err := loadBeatmap()
	if err != nil {
		logMessage(FATAL, "Failed to load beatmap:"+err.Error())
	}
	return &GoRythm{
		movesO:       make(chan [2]int, 2),
		movesX:       make(chan [2]int, 2),
		toBeRemovedO: noMove,
		toBeRemovedX: noMove,
		beatMap:      bm,
	}
}

// Update the game state and return the coordinates where a symbol should be removed or highlighted.
// A maximum of three symbols per player can be placed on the board. When the third symbol is placed, the first symbol is removed.
// The next symbol to be removed in the next round is highlighted.
func (g *GoRythm) Update(playing string, x, y int) (remove, highlight bool, toRemove, toHighlight [2]int) {
	// Check if there is a move to remove
	remove, toRemove = g.moveToRemove(playing)

	// Update the moves
	if playing == "X" {
		if len(g.movesX) == 2 {
			g.toBeRemovedX = <-g.movesX
		}
		g.movesX <- [2]int{x, y}
	} else if playing == "O" {
		if len(g.movesO) == 2 {
			g.toBeRemovedO = <-g.movesO
		}
		g.movesO <- [2]int{x, y}
	} else {
		panic("Invalid player")
	}

	// Check if there is a move to highlight
	highlight, toHighlight = g.moveToHighlight(playing)

	return remove, highlight, toRemove, toHighlight
}

func (g *GoRythm) moveToRemove(playing string) (remove bool, toRemove [2]int) {
	if playing == "X" {
		if g.toBeRemovedX != noMove {
			return true, g.toBeRemovedX
		}
		return false, noMove
	} else if playing == "O" {
		if g.toBeRemovedO != noMove {
			return true, g.toBeRemovedO
		}
		return false, noMove
	}
	panic("Invalid player")
}

func (g *GoRythm) moveToHighlight(playing string) (highlight bool, toHighlight [2]int) {
	if playing == "X" {
		if len(g.movesX) == 2 && g.toBeRemovedX != noMove {
			return true, g.toBeRemovedX
		}
		return false, noMove
	} else if playing == "O" {
		if len(g.movesO) == 2 && g.toBeRemovedO != noMove {
			return true, g.toBeRemovedO
		}
		return false, noMove
	}
	panic("Invalid player")
}

// Calculate the score based on the precision when hitting a beat.
func (g *Game) CalculateScore(beatTime float64) int {
	elapsed := time.Since(g.startTime).Seconds()
	if math.Abs(beatTime-elapsed) < perfectPrec {
		return perfectScore
	} else if math.Abs(beatTime-elapsed) < goodPrec {
		return goodScore
	} else if math.Abs(beatTime-elapsed) < okPrec {
		return okScore
	} else {
		return missedScore
	}
}
