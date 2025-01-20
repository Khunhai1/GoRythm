// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"testing"
)

const (
	sWidth  = 640
	sHeight = 480
)

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g == nil {
		t.Fatal("Expected new game instance, got nil")
	}
	if g.state != StateMenu {
		t.Errorf("Expected initial state to be StateMenu, got %v", g.state)
	}
}

func TestGame_Init(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if g.sWidth != sWidth || g.sHeight != sHeight {
		t.Errorf("Expected screen size to be 640x480, got %dx%d", g.sWidth, g.sHeight)
	}
	if g.audioPlayer == nil {
		t.Error("Expected audio player to be initialized, got nil")
	}
}

func TestGame_Layout(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	screenWidth, screenHeight := g.Layout(800, 600)
	if screenWidth != sWidth || screenHeight != sHeight {
		t.Errorf("Expected layout size to be 640x480, got %dx%d", screenWidth, screenHeight)
	}
}

func TestGame_handleStateMenu(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.handleStateMenu()
	if g.state != StateMenu {
		t.Errorf("Expected state to be StateMenu, got %v", g.state)
	}
}

func TestGame_handleStatePlaying(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.state = StatePlaying
	g.gameMode = EASY_AI_MODE
	g.currentPlayerType = AI_TYPE
	if err := g.handleStatePlaying(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGame_handleStateGameOver(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.state = StateGameOver
	if err := g.handleStateGameOver(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if g.state != StateGameOver {
		t.Errorf("Expected state to be StateGameOver, got %v", g.state)
	}
}

func TestGame_restartGame(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.restartGame()
	if g.rounds != 0 {
		t.Errorf("Expected rounds to be 0, got %d", g.rounds)
	}
	if g.win != NONE_PLAYING {
		t.Errorf("Expected win to be empty, got %s", g.win)
	}
	if g.gameMode != NO_MODE {
		t.Errorf("Expected game mode to be 0, got %d", g.gameMode)
	}
	if g.pointsO != 0 {
		t.Errorf("Expected pointsO to be 0, got %d", g.pointsO)
	}
	if g.pointsX != 0 {
		t.Errorf("Expected pointsX to be 0, got %d", g.pointsX)
	}
}

func TestGame_randomizeStartingPlayer(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.randomizeStartingPlayer()
	if g.currentPlayerSymbol != O_PLAYING && g.currentPlayerSymbol != X_PLAYING {
		t.Errorf("Expected playing to be 'O' or 'X', got %s", g.currentPlayerSymbol)
	}
}

func TestGame_performMove(t *testing.T) {
	g := NewGame()
	if err := g.Init(audioContext, sWidth, sHeight); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	g.randomizeStartingPlayer()
	startingPlayer := g.currentPlayerSymbol
	g.performMove(0, 0)

	if g.board[0][0] != startingPlayer {
		t.Errorf("Expected %s at position (0,0), got %s", startingPlayer, g.board[0][0])
	}
	if g.currentPlayerSymbol == startingPlayer {
		t.Errorf("Expected player to switch, but it didn't")
	}
	if g.rounds != 1 {
		t.Errorf("Expected rounds to be 1, got %d", g.rounds)
	}
}
