package main

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
}

func NewGoRythmGame() *GoRythm {
	return &GoRythm{
		movesO:       make(chan [2]int, 2),
		movesX:       make(chan [2]int, 2),
		toBeRemovedO: noMove,
		toBeRemovedX: noMove,
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
