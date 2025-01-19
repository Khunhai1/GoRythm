package game

import (
	"testing"
)

func TestResetPoints(t *testing.T) {
	g := NewGame()
	g.pointsO = 5
	g.pointsX = 3
	g.ResetPoints()
	if g.pointsO != 0 || g.pointsX != 0 {
		t.Errorf("ResetPoints failed, got pointsO=%d, pointsX=%d", g.pointsO, g.pointsX)
	}
}

func TestPlaceSymbol(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.playing = "O"
	g.placeSymbol(0, 0)
	if g.board[0][0] != "O" {
		t.Errorf("Expected board[0][0] to be O, got %s", g.board[0][0])
	}
}

func TestRemoveSymbol(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.board[1][1] = "O"
	g.removeSymbol(1, 1)

	if g.board[1][1] != "" {
		t.Errorf("removeSymbol failed, expected empty, got %s", g.board[1][1])
	}
}

func TestSwitchPlayer(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.playing = "X"
	g.switchPlayer()
	if g.playing != "O" {
		t.Errorf("switchPlayer failed, expected O, got %s", g.playing)
	}

	g.switchPlayer()
	if g.playing != "X" {
		t.Errorf("switchPlayer failed, expected X, got %s", g.playing)
	}
}

func TestEasyCpu(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.board = [3][3]string{
		{"", "", "X"},
		{"O", "X", "O"},
		{"", "", ""},
	}
	x, y := g.EasyCpu()
	if g.board[x][y] != "" {
		t.Errorf("EasyCpu failed, expected empty cell, got non-empty at (%d, %d)", x, y)
	}
}

func TestCheckWin(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.board = [3][3]string{
		{"X", "X", "X"},
		{"O", "O", ""},
		{"", "", ""},
	}

	winner, _ := g.CheckWin()
	if winner != "X" {
		t.Errorf("CheckWin failed, expected X, got %s", winner)
	}
}

func TestIsBoardFull(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.board = [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	if !g.IsBoardFull() {
		t.Errorf("IsBoardFull failed, expected true, got false")
	}

	g.board[0][0] = ""
	if g.IsBoardFull() {
		t.Errorf("IsBoardFull failed, expected false, got true")
	}
}

func TestHardCpu(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	g.board = [3][3]string{
		{"X", "", "O"},
		{"O", "X", ""},
		{"", "", ""},
	}
	x, y := g.HardCpu()
	if g.board[x][y] != "" {
		t.Errorf("HardCpu failed, expected empty cell, got non-empty at (%d, %d)", x, y)
	}
}

func TestMinimax_WinForX(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)

	// X is about to win
	g.board = [3][3]string{
		{"X", "X", ""},
		{"O", "O", ""},
		{"", "", ""},
	}

	score := g.minimax(0, true)
	if score != 9 { // AI (X) should win
		t.Errorf("expected score 9, got %d", score)
	}
}

func TestMinimax_WinForO(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	// O is about to win
	g.board = [3][3]string{
		{"O", "O", ""},
		{"X", "X", ""},
		{"", "", ""},
	}

	score := g.minimax(0, false)
	if score != -9 { // Player (O) should win
		t.Errorf("expected score -9, got %d", score)
	}
}

func TestMinimax_Draw(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	// The board is full, and it's a draw
	g.board = [3][3]string{
		{"X", "O", "X"},
		{"X", "X", "O"},
		{"O", "X", "O"},
	}

	score := g.minimax(0, true)
	if score != 0 { // Draw should return a score of 0
		t.Errorf("expected score 0 for draw, got %d", score)
	}
}

func TestMinimax_BestMove(t *testing.T) {
	g := NewGame()
	g.Init(audioContext, 480, 700)
	// The AI (X) has multiple moves to choose from
	g.board = [3][3]string{
		{"X", "O", ""},
		{"", "X", "O"},
		{"O", "", ""},
	}

	score := g.minimax(0, true)
	if score != 9 { // AI (X) should find the winning move
		t.Errorf("expected score 9, got %d", score)
	}
}
