// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"testing"
)

// TestPlaceSymbol tests the placeSymbol function.
// Checks if the symbol is placed correctly on the board.
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

// TestRemoveSymbol tests the removeSymbol function.
// Checks if the symbol is removed correctly from the board.
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

// TestSwitchPlayer tests the switchPlayer function.
// Checks if the current player is switched correctly based on the current player.
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

// TestEasyCpu tests the EasyCpu function.
// Checks if the function returns a valid move.
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

// TestcheckWinBoard tests the checkWinBoard function.
// Checks if the function returns the correct winner based on the board.
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

	winner, _ := g.checkWinBoard()
	if winner != X_PLAYING {
		t.Errorf("checkWinBoard failed, expected X, got %s", winner)
	}
}

// FuzzcheckWinBoard is a fuzzing test for the checkWinBoard function.
// It generates random board configurations and checks if the function returns a valid winner.
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
		winner, _ := g.checkWinBoard()
		if winner != NONE_PLAYING && winner != X_PLAYING && winner != O_PLAYING {
			t.Errorf("checkWinBoard returned an invalid winner: %s", winner)
		}
	})
}

// TestIsBoardFull tests the isBoardFull function.
// Checks if the function returns true when the board is full and false otherwise.
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

	if !g.isBoardFull() {
		t.Errorf("isBoardFull failed, expected true, got false")
	}

	g.board[0][0] = NONE_PLAYING
	if g.isBoardFull() {
		t.Errorf("isBoardFull failed, expected false, got true")
	}
}

// TestHardCpu tests the HardCpu function.
// Checks if the function returns a valid move.
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

// TestMinimax tests the minimax function for the player X.
// Checks if the function returns the correct score for the current board state,
// the player X should always win.
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

// TestMinimax tests the minimax function for the player O.
// Checks if the function returns the correct score for the current board state,
// the player O should always win.
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

// TestMinimax_Draw tests the minimax function for a draw.
// Checks if the function returns the correct score for the current board state,
// the game should end in a draw.
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

// TestMinimax_BestMove tests the minimax function for the best move.
// Checks if the function returns the correct score for the current board state,
// the AI (X) should find the best move to win.
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

// TestCheckWinScore tests the checkWinScore function.
// Checks if the function returns the correct winner based on the score.
func TestCheckWinScore(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	winner := g.checkWinScore()

	if winner != NONE_PLAYING {
		t.Errorf("checkWinScore failed, expected no winner, got %s", winner)
	}

	g.pointsX = 500
	g.pointsO = 300

	winner = g.checkWinScore()
	if winner != X_PLAYING {
		t.Errorf("checkWinScore failed, expected X, got %s", winner)
	}

	g.pointsO += 400
	winner = g.checkWinScore()
	if winner != O_PLAYING {
		t.Errorf("checkWinScore failed, expected O, got %s", winner)
	}
}
