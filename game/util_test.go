// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

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
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.currentPlayerSymbol = O_PLAYING
	g.placeSymbol(0, 0)
	if g.board[0][0] != O_PLAYING {
		t.Errorf("Expected board[0][0] to be O, got %s", g.board[0][0])
	}
}

func TestRemoveSymbol(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.board[1][1] = O_PLAYING
	g.removeSymbol(1, 1)

	if g.board[1][1] != NONE_PLAYING {
		t.Errorf("removeSymbol failed, expected empty, got %s", g.board[1][1])
	}
}

func TestSwitchPlayer(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.currentPlayerSymbol = X_PLAYING
	g.switchPlayer()
	if g.currentPlayerSymbol != O_PLAYING {
		t.Errorf("switchPlayer failed, expected O, got %s", g.currentPlayerSymbol)
	}

	g.switchPlayer()
	if g.currentPlayerSymbol != X_PLAYING {
		t.Errorf("switchPlayer failed, expected X, got %s", g.currentPlayerSymbol)
	}
}

func TestEasyCpu(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.board = [3][3]SymbolPlaying{
		{NONE_PLAYING, NONE_PLAYING, X_PLAYING},
		{O_PLAYING, X_PLAYING, O_PLAYING},
		{NONE_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}
	x, y := g.EasyCpu()
	if g.board[x][y] != NONE_PLAYING {
		t.Errorf("EasyCpu failed, expected empty cell, got non-empty at (%d, %d)", x, y)
	}
}

func TestCheckWinBoard(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, X_PLAYING, X_PLAYING},
		{O_PLAYING, O_PLAYING, NONE_PLAYING},
		{NONE_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}

	winner, _ := g.CheckWinBoard()
	if winner != X_PLAYING {
		t.Errorf("CheckWinBoard failed, expected X, got %s", winner)
	}
}

func FuzzCheckWinBoard(f *testing.F) {
	// Add some seed cases (optional)
	f.Add("XXX------")
	f.Add("O--O--O--")
	f.Add("XOXOXOXOX")

	f.Fuzz(func(t *testing.T, board string) {
		if len(board) != 9 {
			t.Skip("Invalid board configuration")
		}

		// Convert the string back to a 3x3 array
		var arrBoard [3][3]SymbolPlaying
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cell := SymbolPlaying(board[i*3+j])
				if cell == "-" {
					cell = NONE_PLAYING
				}
				if cell != X_PLAYING && cell != O_PLAYING && cell != NONE_PLAYING {
					t.Skip("Invalid board configuration")
				}
				arrBoard[i][j] = cell
			}
		}

		// Create a game instance and set the board
		g := NewGame()
		g.board = arrBoard

		// Check for a winner
		winner, _ := g.CheckWinBoard()
		if winner != NONE_PLAYING && winner != X_PLAYING && winner != O_PLAYING {
			t.Errorf("CheckWinBoard returned an invalid winner: %s", winner)
		}
	})
}

func TestIsBoardFull(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, O_PLAYING, X_PLAYING},
		{O_PLAYING, X_PLAYING, O_PLAYING},
		{X_PLAYING, O_PLAYING, X_PLAYING},
	}

	if !g.IsBoardFull() {
		t.Errorf("IsBoardFull failed, expected true, got false")
	}

	g.board[0][0] = NONE_PLAYING
	if g.IsBoardFull() {
		t.Errorf("IsBoardFull failed, expected false, got true")
	}
}

func TestHardCpu(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, NONE_PLAYING, O_PLAYING},
		{O_PLAYING, X_PLAYING, NONE_PLAYING},
		{NONE_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}
	x, y := g.HardCpu()
	if g.board[x][y] != NONE_PLAYING {
		t.Errorf("HardCpu failed, expected empty cell, got non-empty at (%d, %d)", x, y)
	}
}

func TestMinimax_WinForX(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// X is about to win
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, X_PLAYING, NONE_PLAYING},
		{O_PLAYING, O_PLAYING, NONE_PLAYING},
		{NONE_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}

	score := g.minimax(0, true)
	if score != 9 { // AI (X) should win
		t.Errorf("expected score 9, got %d", score)
	}
}

func TestMinimax_WinForO(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// O is about to win
	g.board = [3][3]SymbolPlaying{
		{O_PLAYING, O_PLAYING, NONE_PLAYING},
		{X_PLAYING, X_PLAYING, NONE_PLAYING},
		{NONE_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}

	score := g.minimax(0, false)
	if score != -9 { // Player (O) should win
		t.Errorf("expected score -9, got %d", score)
	}
}

func TestMinimax_Draw(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// The board is full, and it's a draw
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, O_PLAYING, X_PLAYING},
		{X_PLAYING, X_PLAYING, O_PLAYING},
		{O_PLAYING, X_PLAYING, O_PLAYING},
	}

	score := g.minimax(0, true)
	if score != 0 { // Draw should return a score of 0
		t.Errorf("expected score 0 for draw, got %d", score)
	}
}

func TestMinimax_BestMove(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// The AI (X) has multiple moves to choose from
	g.board = [3][3]SymbolPlaying{
		{X_PLAYING, O_PLAYING, NONE_PLAYING},
		{NONE_PLAYING, X_PLAYING, O_PLAYING},
		{O_PLAYING, NONE_PLAYING, NONE_PLAYING},
	}

	score := g.minimax(0, true)
	if score != 9 { // AI (X) should find the winning move
		t.Errorf("expected score 9, got %d", score)
	}
}
