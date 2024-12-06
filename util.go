package main

import (
	"math/rand"
	"time"
)

func newRandom() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}

func (g *Game) CheckWin() string {
	for i, _ := range g.gameBoard {
		if g.gameBoard[i][0] == g.gameBoard[i][1] && g.gameBoard[i][1] == g.gameBoard[i][2] {
			return g.gameBoard[i][0]
		}
	}
	for i, _ := range g.gameBoard {
		if g.gameBoard[0][i] == g.gameBoard[1][i] && g.gameBoard[1][i] == g.gameBoard[2][i] {
			return g.gameBoard[0][i]
		}
	}
	if (g.gameBoard[0][0] == g.gameBoard[1][1] && g.gameBoard[1][1] == g.gameBoard[2][2]) || (g.gameBoard[0][2] == g.gameBoard[1][1] && g.gameBoard[1][1] == g.gameBoard[2][0]) {
		return g.gameBoard[1][1]
	}
	if g.round == 8 {
		return "tie"
	}
	return ""
}

func (g *Game) wins(winner string) {
	if winner == "O" {
		g.win = "O"
		g.pointsO++
		g.state = 2
	} else if winner == "X" {
		g.win = "X"
		g.pointsX++
		g.state = 2
	} else if winner == "tie" {
		g.win = "No one\n"
		g.state = 2
	}
}

func (g *Game) ResetPoints() {
	g.pointsO = 0
	g.pointsX = 0
}
